import { experiences } from "@/data/portfolio";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { useInView } from "@/hooks/useInView";

export default function ExperienceSection() {
  const { ref, isInView } = useInView(0.1);

  return (
    <section id="experience" className="py-24 px-6 bg-secondary/30">
      <div ref={ref} className="max-w-6xl mx-auto">
        <h2
          className={`text-3xl md:text-4xl font-bold mb-2 section-title ${isInView ? "animate-fade-in-up" : "opacity-0"
            }`}
        >
          Experience
        </h2>
        <p
          className={`text-muted-foreground mb-12 mt-6 ${isInView ? "animate-fade-in-up animate-delay-100" : "opacity-0"
            }`}
        >
          Professional experience and career journey.
        </p>

        <div className="relative pl-12">
          {/* Timeline line */}
          <div className="timeline-line" />

          <div className="space-y-8">
            {experiences.map((exp, index) => (
              <div
                key={index}
                className={`relative ${isInView ? "animate-fade-in-left" : "opacity-0"
                  }`}
                style={{ animationDelay: `${(index + 2) * 0.1}s` }}
              >
                {/* Timeline dot */}
                <div className="timeline-dot" />

                <Card className="card-hover bg-card/50 border-border/50">
                  <CardContent className="p-6">
                    <div className="flex flex-col sm:flex-row sm:items-start sm:justify-between mb-2">
                      <div>
                        <h3 className="text-lg font-semibold text-foreground">
                          {exp.position}
                        </h3>
                        <p className="text-base font-medium text-primary">
                          {exp.company}
                        </p>
                      </div>
                      <span className="text-sm text-muted-foreground font-medium mt-1 sm:mt-0 sm:text-right whitespace-nowrap">
                        {exp.startDate} — {exp.endDate}
                      </span>
                    </div>

                    <p className="text-sm text-muted-foreground/80 mt-3 leading-relaxed">
                      {exp.description}
                    </p>

                    <div className="flex flex-wrap gap-2 mt-4">
                      {exp.technologies.map((tech) => (
                        <Badge
                          key={tech}
                          variant="secondary"
                          className="text-xs font-medium bg-primary/10 text-primary border-primary/20 hover:bg-primary/20"
                        >
                          {tech}
                        </Badge>
                      ))}
                    </div>
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
