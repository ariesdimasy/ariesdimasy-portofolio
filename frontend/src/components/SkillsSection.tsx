import { skills } from "@/data/portfolio";
import { useInView } from "@/hooks/useInView";

const categoryIcons: Record<string, string> = {
  Frontend: "🎨",
  Backend: "⚙️",
  Database: "🗃️",
  Tools: "🛠️",
};

const categoryGradients: Record<string, string> = {
  Frontend: "from-blue-500/20 to-cyan-500/20",
  Backend: "from-violet-500/20 to-purple-500/20",
  Database: "from-emerald-500/20 to-green-500/20",
  Tools: "from-orange-500/20 to-amber-500/20",
};

export default function SkillsSection() {
  const { ref, isInView } = useInView(0.1);

  const categories = [...new Set(skills.map((s) => s.category))];

  return (
    <section id="skills" className="py-24 px-6 bg-secondary/30">
      <div ref={ref} className="max-w-6xl mx-auto">
        <h2
          className={`text-3xl md:text-4xl font-bold mb-2 section-title ${
            isInView ? "animate-fade-in-up" : "opacity-0"
          }`}
        >
          Skills
        </h2>
        <p
          className={`text-muted-foreground mb-12 mt-6 ${
            isInView ? "animate-fade-in-up animate-delay-100" : "opacity-0"
          }`}
        >
          Technologies and tools I work with.
        </p>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          {categories.map((category, catIndex) => (
            <div
              key={category}
              className={`${isInView ? "animate-fade-in-up" : "opacity-0"}`}
              style={{ animationDelay: `${(catIndex + 2) * 0.1}s` }}
            >
              <div className="flex items-center gap-3 mb-5">
                <span className="text-2xl">{categoryIcons[category]}</span>
                <h3 className="text-lg font-semibold text-foreground">
                  {category}
                </h3>
              </div>
              <div className="flex flex-wrap gap-3">
                {skills
                  .filter((s) => s.category === category)
                  .map((skill, skillIndex) => (
                    <div
                      key={skill.name}
                      className={`px-4 py-2.5 rounded-xl bg-gradient-to-r ${categoryGradients[category]} border border-border/30 text-sm font-medium text-foreground hover:scale-105 transition-all duration-200 cursor-default hover:border-primary/30 hover:shadow-lg hover:shadow-primary/5`}
                      style={{
                        animationDelay: `${(catIndex * 5 + skillIndex + 3) * 0.05}s`,
                      }}
                    >
                      {skill.name}
                    </div>
                  ))}
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
