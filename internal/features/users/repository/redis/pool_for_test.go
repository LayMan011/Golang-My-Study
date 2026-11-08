package users_redis_repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"time"

	core_redis_pool "github.com/LayMan011/Golang-My-Study/internal/core/repository/redis/pool"
	"github.com/redis/go-redis/v9"
)

type PoolStub struct {
	mu     sync.RWMutex
	values map[string]string
	hashes map[string]map[string]string
	closed bool
}

func NewPoolStub() *PoolStub {
	return &PoolStub{
		values: make(map[string]string),
		hashes: make(map[string]map[string]string),
	}
}

func (p *PoolStub) Set(
	ctx context.Context,
	key string,
	value any,
	_ time.Duration,
) core_redis_pool.CustomStatusCmd {
	if err := ctx.Err(); err != nil {
		return statusCmd{err: err}
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return statusCmd{err: errors.New("redis stub is closed")}
	}

	p.values[key] = fmt.Sprint(value)
	return statusCmd{}
}

func (p *PoolStub) Get(ctx context.Context, key string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return "", errors.New("redis stub is closed")
	}

	value, ok := p.values[key]
	if !ok {
		return "", errors.New("redis: nil")
	}
	return value, nil
}

func (p *PoolStub) DoError(
	ctx context.Context,
	_ string,
	_ int,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return errors.New("redis stub is closed")
	}
	return nil
}

func (p *PoolStub) HSet(
	ctx context.Context,
	key string,
	value any,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	fields, err := hashFields(value)
	if err != nil {
		return err
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return errors.New("redis stub is closed")
	}

	if p.hashes[key] == nil {
		p.hashes[key] = make(map[string]string)
	}
	for field, fieldValue := range fields {
		p.hashes[key][field] = fieldValue
	}
	return nil
}

func (p *PoolStub) HGetAll(
	ctx context.Context,
	key string,
) (map[string]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return nil, errors.New("redis stub is closed")
	}

	result := make(map[string]string, len(p.hashes[key]))
	for field, value := range p.hashes[key] {
		result[field] = value
	}
	return result, nil
}

func (p *PoolStub) Expire(
	ctx context.Context,
	_ string,
	_ time.Duration,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return errors.New("redis stub is closed")
	}
	// TTL специально не эмулируется.
	return nil
}

func (p *PoolStub) Pipelined(
	ctx context.Context,
	fn func(core_redis_pool.Pipeline) error,
) (core_redis_pool.CustomCmder, error) {
	if err := ctx.Err(); err != nil {
		return statusCmd{err: err}, err
	}

	pipe := &pipelineStub{pool: p, ctx: ctx}
	if err := fn(pipe); err != nil {
		return statusCmd{err: err}, err
	}
	if err := pipe.err; err != nil {
		return statusCmd{err: err}, err
	}
	return statusCmd{}, nil
}

func (p *PoolStub) Exists(
	ctx context.Context,
	key string,
) core_redis_pool.CustomIntCmd {
	if err := ctx.Err(); err != nil {
		return intCmd{err: err}
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return intCmd{err: errors.New("redis stub is closed")}
	}

	if _, ok := p.values[key]; ok {
		return intCmd{value: 1}
	}
	if _, ok := p.hashes[key]; ok {
		return intCmd{value: 1}
	}
	return intCmd{}
}

func (p *PoolStub) HDel(
	ctx context.Context,
	key string,
	fields ...string,
) core_redis_pool.CustomIntCmd {
	if err := ctx.Err(); err != nil {
		return intCmd{err: err}
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return intCmd{err: errors.New("redis stub is closed")}
	}

	deleted := int64(0)
	for _, field := range fields {
		if _, ok := p.hashes[key][field]; ok {
			delete(p.hashes[key], field)
			deleted++
		}
	}
	return intCmd{value: deleted}
}

func (p *PoolStub) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.closed = true
	return nil
}

func (p *PoolStub) Incr(ctx context.Context, key string) (int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return 0, errors.New("redis stub is closed")
	}

	value, err := strconv.ParseInt(p.values[key], 10, 64)
	if err != nil && p.values[key] != "" {
		return 0, err
	}

	value++
	p.values[key] = strconv.FormatInt(value, 10)
	return value, nil
}

func (p *PoolStub) Del(_ context.Context, keys ...string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, key := range keys {
		delete(p.hashes, key)
	}

	return nil
}

type pipelineStub struct {
	pool *PoolStub
	ctx  context.Context
	err  error
}

func (pipe *pipelineStub) HSet(
	ctx context.Context,
	key string,
	field string,
	value any,
) error {
	if pipe.err != nil {
		return pipe.err
	}
	if err := ctx.Err(); err != nil {
		pipe.err = err
		return err
	}

	pipe.pool.mu.Lock()
	defer pipe.pool.mu.Unlock()

	if pipe.pool.closed {
		pipe.err = errors.New("redis stub is closed")
		return pipe.err
	}
	if pipe.pool.hashes[key] == nil {
		pipe.pool.hashes[key] = make(map[string]string)
	}
	pipe.pool.hashes[key][field] = fmt.Sprint(value)
	return nil
}

func (pipe *pipelineStub) Expire(
	ctx context.Context,
	_ string,
	_ time.Duration,
) error {
	if pipe.err != nil {
		return pipe.err
	}
	if err := ctx.Err(); err != nil {
		pipe.err = err
		return err
	}

	// TTL специально не эмулируется.
	return nil
}

type statusCmd struct {
	err error
}

func (cmd statusCmd) Err() error {
	return cmd.err
}

type intCmd struct {
	value int64
	err   error
}

func (cmd intCmd) Result() (int64, error) {
	return cmd.value, cmd.err
}

func hashFields(value any) (map[string]string, error) {
	result := make(map[string]string)

	if value == nil {
		return nil, errors.New("cannot HSet a nil value")
	}

	if values, ok := value.(map[string]string); ok {
		for key, val := range values {
			result[key] = val
		}
		return result, nil
	}

	if values, ok := value.(map[string]any); ok {
		for key, val := range values {
			result[key] = fmt.Sprint(val)
		}
		return result, nil
	}

	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil, errors.New("cannot HSet a nil pointer")
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("unsupported HSet value type %T", value)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("redis")
		if name == "" || name == "-" {
			continue
		}
		if !v.Field(i).CanInterface() {
			continue
		}
		result[name] = stringify(v.Field(i).Interface())
	}

	return result, nil
}

func stringify(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprint(value)
	}
}

type recordingPool struct {
	core_redis_pool.Pool
	values map[string]string
}

func (p *recordingPool) Set(
	_ context.Context,
	key string,
	value any,
	_ time.Duration,
) core_redis_pool.CustomStatusCmd {
	p.values[key] = value.(string)
	return successfulStatusCmd{}
}

type successfulStatusCmd struct{}

func (successfulStatusCmd) Err() error { return nil }

type miniredisPool struct {
	*PoolStub
	client *redis.Client
}

var _ core_redis_pool.Pool = (*miniredisPool)(nil)

func (p *miniredisPool) Set(
	ctx context.Context,
	key string,
	value any,
	expiration time.Duration,
) core_redis_pool.CustomStatusCmd {
	return p.client.Set(ctx, key, value, expiration)
}

func (p *miniredisPool) Incr(ctx context.Context, key string) (int64, error) {
	return p.client.Incr(ctx, key).Result()
}

func (p *miniredisPool) Del(ctx context.Context, keys ...string) error {
	return p.client.Del(ctx, keys...).Err()
}

func (p *miniredisPool) Del(ctx context.Context, keys ...string) error {
	return p.client.Del(ctx, keys...).Err()
}
