import { biodata } from "@/data/portfolio";
import { Button } from "@/components/ui/button";
import { useInView } from "@/hooks/useInView";
import ParticleBackground from "@/components/ParticleBackground";
import myPhoto from "@/assets/myphoto.png";

export default function HeroSection() {
  const { ref, isInView } = useInView(0.2);

  return (
    <section
      id="home"
      className="min-h-screen flex items-center justify-center hero-gradient relative overflow-hidden"
    >
      {/* Particle Background */}
      <ParticleBackground />

      {/* Subtle glow orbs */}
      <div className="absolute top-20 left-10 w-72 h-72 rounded-full bg-primary/5 blur-3xl pointer-events-none" />
      <div className="absolute bottom-20 right-10 w-96 h-96 rounded-full bg-accent/5 blur-3xl pointer-events-none" />

      {/* Main content: two-column layout */}
      <div
        ref={ref}
        className="max-w-6xl mx-auto px-6 relative z-10 flex flex-col-reverse md:flex-row items-center justify-between gap-12 py-20"
      >
        {/* ── LEFT: Text content ── */}
        <div className="flex-1 text-center md:text-left">
          {/* Greeting badge */}
          <div
            className={`inline-flex items-center gap-2 px-4 py-2 rounded-full border border-primary/30 bg-primary/10 mb-8 ${isInView ? "animate-fade-in-up" : "opacity-0"
              }`}
          >
            <span className="w-2 h-2 rounded-full bg-green-400 animate-pulse" />
            <span className="text-sm text-primary font-medium">Available for work</span>
          </div>

          {/* Name */}
          <h1
            className={`text-5xl md:text-6xl lg:text-7xl font-extrabold mb-6 tracking-tight ${isInView ? "animate-fade-in-up animate-delay-100" : "opacity-0"
              }`}
          >
            Hi, I'm{" "}
            <span className="gradient-text">{biodata.name}</span>
          </h1>

          {/* Tagline */}
          <p
            className={`text-xl md:text-2xl text-muted-foreground mb-4 font-medium ${isInView ? "animate-fade-in-up animate-delay-200" : "opacity-0"
              }`}
          >
            {biodata.tagline}
          </p>

          {/* Description */}
          <p
            className={`text-base md:text-lg text-muted-foreground/80 max-w-xl mb-10 leading-relaxed ${isInView ? "animate-fade-in-up animate-delay-300" : "opacity-0"
              }`}
          >
            {biodata.description}
          </p>

          {/* CTA Buttons */}
          <div
            className={`flex flex-wrap items-center justify-center md:justify-start gap-4 ${isInView ? "animate-fade-in-up animate-delay-400" : "opacity-0"
              }`}
          >
            <Button
              size="lg"
              className="bg-gradient-to-r from-primary to-accent hover:opacity-90 text-white font-semibold px-8 py-6 text-base cursor-pointer"
              onClick={() =>
                document.querySelector("#projects")?.scrollIntoView({ behavior: "smooth" })
              }
            >
              View My Work
            </Button>
            <a
              href={`mailto:${biodata.email}`}
              className="inline-flex items-center justify-center rounded-lg border border-primary/40 text-primary hover:bg-primary/10 font-semibold px-8 py-3 text-base transition-all"
            >
              Contact Me
            </a>
          </div>

          {/* Social links */}
          <div
            className={`flex items-center justify-center md:justify-start gap-6 mt-10 ${isInView ? "animate-fade-in-up animate-delay-500" : "opacity-0"
              }`}
          >
            <a
              href={biodata.github}
              target="_blank"
              rel="noopener noreferrer"
              className="text-muted-foreground hover:text-primary transition-colors"
              aria-label="GitHub"
            >
              <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
              </svg>
            </a>
            <a
              href={biodata.linkedin}
              target="_blank"
              rel="noopener noreferrer"
              className="text-muted-foreground hover:text-primary transition-colors"
              aria-label="LinkedIn"
            >
              <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z" />
              </svg>
            </a>
            <a
              href={`mailto:${biodata.email}`}
              className="text-muted-foreground hover:text-primary transition-colors"
              aria-label="Email"
            >
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={1.5}>
                <path strokeLinecap="round" strokeLinejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75" />
              </svg>
            </a>
          </div>
        </div>

        {/* ── RIGHT: Profile photo ── */}
        <div
          className={` ${isInView ? "animate-fade-in-up animate-delay-200" : "opacity-0"
            }`}

        >
          {/* Outer glow ring */}
          <div className="relative">


            {/* Photo container */}
            <div
              className="relative w-64 h-64 md:w-120 md:h-120 rounded-full overflow-hidden"
              style={{
                zIndex: 1,

                border: "2px solid oklch(0.65 0.2 270 / 0.4)",
              }}
            >
              <img
                src={myPhoto}
                alt={biodata.name}
                className="absolute bottom-0 left-1/2 -translate-x-1/2 h-[105%] w-auto object-contain object-bottom"
                style={{ mixBlendMode: "luminosity", filter: "contrast(1.1) brightness(0.95)" }}
              />
            </div>

            {/* Decorative dots */}
            <div className="absolute -top-3 -right-3 w-5 h-5 rounded-full bg-primary animate-pulse" />
            <div className="absolute -bottom-2 -left-4 w-3 h-3 rounded-full bg-accent animate-pulse" style={{ animationDelay: "0.5s" }} />
          </div>
        </div>
      </div>

      {/* Scroll indicator */}
      <div className="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center gap-2 animate-bounce">
        <span className="text-xs text-muted-foreground">Scroll</span>
        <svg className="w-4 h-4 text-muted-foreground" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 14l-7 7m0 0l-7-7m7 7V3" />
        </svg>
      </div>
    </section>
  );
}
