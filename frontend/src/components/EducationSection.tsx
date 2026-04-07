import { educations } from "@/data/portfolio";
import { Card, CardContent } from "@/components/ui/card";
import { useInView } from "@/hooks/useInView";

export default function EducationSection() {
  const { ref, isInView } = useInView(0.1);

  return (
    <section id="education" className="py-24 px-6">
      <div ref={ref} className="max-w-6xl mx-auto">
        <h2
          className={`text-3xl md:text-4xl font-bold mb-2 section-title ${isInView ? "animate-fade-in-up" : "opacity-0"
            }`}
        >
          Education
        </h2>
        <p
          className={`text-muted-foreground mb-12 mt-6 ${isInView ? "animate-fade-in-up animate-delay-100" : "opacity-0"
            }`}
        >
          My academic journey and educational background.
        </p>

        <div className="relative pl-12">
          {/* Timeline line */}
          <div className="timeline-line" />

          <div className="space-y-8">
            {educations.map((edu, index) => (
              <div
                key={index}
                className={`relative ${isInView
                  ? `animate-fade-in-left animate-delay-${(index + 2) * 100}`
                  : "opacity-0"
                  }`}
                style={{ animationDelay: `${(index + 2) * 0.1}s` }}
              >
                {/* Timeline dot */}
                <div className="timeline-dot" />

                <Card className="card-hover bg-card/50 border-border/50">
                  <CardContent className="p-6">
                    <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-2">
                      <h3 className="text-lg font-semibold text-foreground">
                        {edu.institution}
                      </h3>
                      <span className="text-sm text-primary font-medium mt-1 sm:mt-0">
                        {edu.startYear} — {edu.endYear}
                      </span>
                    </div>
                    <p className="text-base font-medium text-muted-foreground mb-1">
                      {edu.degree}
                    </p>
                    <p className="text-sm text-accent">{edu.field}</p>
                    {edu.description && (
                      <p className="text-sm text-muted-foreground/80 mt-3 leading-relaxed">
                        {edu.description}
                      </p>
                    )}
                  </CardContent>
                </Card>
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}
