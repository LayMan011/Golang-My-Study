const subjects = [
  { id: 'all', name: 'Все предметы', color: 'primary' },
  { id: 'russian', name: 'Русский язык', color: 'subject-russian' },
  { id: 'english', name: 'Английский язык', color: 'subject-english' },
  { id: 'math', name: 'Математика', color: 'subject-math' },
  { id: 'physics', name: 'Физика', color: 'subject-physics' },
  { id: 'chemistry', name: 'Химия', color: 'subject-chemistry' },
  { id: 'biology', name: 'Биология', color: 'subject-biology' },
  { id: 'history', name: 'История', color: 'subject-history' },
  { id: 'social', name: 'Обществознание', color: 'subject-social' },
];

export const SubjectSerchHelpers = ({selectedSubject, setSelectedSubject} : {selectedSubject: string, setSelectedSubject: React.Dispatch<React.SetStateAction<string>>}) => {
  return (
    <div className="flex flex-wrap gap-2 mt-6">
      {subjects.map(subject => (
        <button
          key={subject.id}
          onClick={() => setSelectedSubject(subject.id)}
          className={`px-4 py-2 rounded-full text-sm font-medium transition-all ${
            selectedSubject === subject.id
              ? 'text-white shadow-md scale-105'
              : 'bg-card text-muted-foreground hover:bg-muted border border-border'
          }`}
          style={{
            backgroundColor: selectedSubject === subject.id ? `var(--${subject.color})` : undefined,
          }}
        >
          {subject.name}
        </button>
      ))}
    </div>
  )
}