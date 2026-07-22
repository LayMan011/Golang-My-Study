export const ProfileCoursesSkeleton = () => {
  const skeletonItems = Array.from({ length: 3 });

  return (
    <section className="py-12">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8 space-y-6">
        {/* Заголовок вкладок */}
        <div className="flex gap-4">
          <div className="h-9 w-24 bg-muted rounded-full animate-pulse" />
          <div className="h-9 w-32 bg-muted rounded-full animate-pulse" />
        </div>

        {/* Карточки курсов */}
        {skeletonItems.map((_, idx) => (
          <div
            key={idx}
            className="bg-card rounded-2xl p-6 border border-border shadow-sm"
          >
            <div className="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-4">
              <div className="flex-1 space-y-3">
                <div className="h-6 w-32 bg-muted rounded-full animate-pulse" />
                <div className="h-5 w-64 bg-muted rounded-md animate-pulse" />
                <div className="h-4 w-40 bg-muted rounded-md animate-pulse" />
              </div>
              <div className="h-10 w-32 bg-muted rounded-lg animate-pulse" />
            </div>

            <div className="space-y-2">
              <div className="flex items-center justify-between text-sm">
                <div className="h-4 w-40 bg-muted rounded-md animate-pulse" />
                <div className="h-4 w-10 bg-muted rounded-md animate-pulse" />
              </div>
              <div className="h-3 w-full bg-muted rounded-full animate-pulse" />
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};