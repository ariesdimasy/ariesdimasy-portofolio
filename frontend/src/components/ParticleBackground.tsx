import { useEffect, useRef } from "react";

interface Particle {
  x: number;
  y: number;
  vx: number;
  vy: number;
  radius: number;
  opacity: number;
  color: string;
  pulsePhase: number;
  pulseSpeed: number;
}

const PARTICLE_COUNT = 160;
const MAX_DISTANCE = 120;
const MOUSE_REPEL_RADIUS = 120;
const MOUSE_REPEL_FORCE = 0.5;
const MIN_SPEED = 0.25;
const MAX_SPEED = 1.2;

const COLORS = [
  "139, 92, 246,",   // violet/primary
  "99, 102, 241,",   // indigo
  "168, 85, 247,",   // purple/accent
  "167, 139, 250,",  // light violet
  "129, 140, 248,",  // light indigo
];

function clampSpeed(vx: number, vy: number): [number, number] {
  const speed = Math.sqrt(vx * vx + vy * vy);
  if (speed === 0) {
    // Give random direction if stalled
    const angle = Math.random() * Math.PI * 2;
    return [Math.cos(angle) * MIN_SPEED, Math.sin(angle) * MIN_SPEED];
  }
  const clamped = Math.max(MIN_SPEED, Math.min(MAX_SPEED, speed));
  return [(vx / speed) * clamped, (vy / speed) * clamped];
}

export default function ParticleBackground() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const particlesRef = useRef<Particle[]>([]);
  const mouseRef = useRef({ x: -9999, y: -9999 });
  const rafRef = useRef<number>(0);
  const frameRef = useRef<number>(0);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext("2d");
    if (!ctx) return;

    const resize = () => {
      canvas.width = canvas.offsetWidth;
      canvas.height = canvas.offsetHeight;
    };

    const init = () => {
      particlesRef.current = [];
      for (let i = 0; i < PARTICLE_COUNT; i++) {
        const angle = Math.random() * Math.PI * 2;
        const speed = MIN_SPEED + Math.random() * (MAX_SPEED - MIN_SPEED);
        particlesRef.current.push({
          x: Math.random() * canvas.width,
          y: Math.random() * canvas.height,
          vx: Math.cos(angle) * speed,
          vy: Math.sin(angle) * speed,
          radius: Math.random() * 2.5 + 0.8,
          opacity: Math.random() * 0.55 + 0.2,
          color: COLORS[Math.floor(Math.random() * COLORS.length)],
          pulsePhase: Math.random() * Math.PI * 2,
          pulseSpeed: 0.01 + Math.random() * 0.02,
        });
      }
    };

    const draw = () => {
      frameRef.current++;
      ctx.clearRect(0, 0, canvas.width, canvas.height);

      const particles = particlesRef.current;
      const mouse = mouseRef.current;
      const t = frameRef.current;

      for (let i = 0; i < particles.length; i++) {
        const p = particles[i];

        // Apply a subtle Perlin-like drift so particles never feel stuck
        // using a simple sinusoidal nudge per particle
        const driftX = Math.sin(t * 0.008 + p.pulsePhase) * 0.04;
        const driftY = Math.cos(t * 0.009 + p.pulsePhase * 1.3) * 0.04;
        p.vx += driftX;
        p.vy += driftY;

        // Mouse repel
        const dx = p.x - mouse.x;
        const dy = p.y - mouse.y;
        const dist = Math.sqrt(dx * dx + dy * dy);
        if (dist < MOUSE_REPEL_RADIUS && dist > 0) {
          const force = ((MOUSE_REPEL_RADIUS - dist) / MOUSE_REPEL_RADIUS) * MOUSE_REPEL_FORCE;
          p.vx += (dx / dist) * force;
          p.vy += (dy / dist) * force;
        }

        // Enforce min/max speed — particles NEVER stop
        [p.vx, p.vy] = clampSpeed(p.vx, p.vy);

        p.x += p.vx;
        p.y += p.vy;

        // Wrap around edges (seamless, more alive than bouncing)
        if (p.x < -5)  p.x = canvas.width + 5;
        if (p.x > canvas.width + 5)  p.x = -5;
        if (p.y < -5)  p.y = canvas.height + 5;
        if (p.y > canvas.height + 5) p.y = -5;

        // Pulsing opacity
        p.pulsePhase += p.pulseSpeed;
        const pulseOpacity = p.opacity * (0.75 + 0.25 * Math.sin(p.pulsePhase));

        // Draw particle with a tiny glow
        const grd = ctx.createRadialGradient(p.x, p.y, 0, p.x, p.y, p.radius * 2.5);
        grd.addColorStop(0, `rgba(${p.color} ${pulseOpacity})`);
        grd.addColorStop(1, `rgba(${p.color} 0)`);
        ctx.beginPath();
        ctx.arc(p.x, p.y, p.radius * 2.5, 0, Math.PI * 2);
        ctx.fillStyle = grd;
        ctx.fill();

        // Solid core
        ctx.beginPath();
        ctx.arc(p.x, p.y, p.radius, 0, Math.PI * 2);
        ctx.fillStyle = `rgba(${p.color} ${pulseOpacity})`;
        ctx.fill();

        // Draw connections
        for (let j = i + 1; j < particles.length; j++) {
          const p2 = particles[j];
          const ex = p.x - p2.x;
          const ey = p.y - p2.y;
          const edist = Math.sqrt(ex * ex + ey * ey);

          if (edist < MAX_DISTANCE) {
            const alpha = (1 - edist / MAX_DISTANCE) * 0.35;
            ctx.beginPath();
            ctx.strokeStyle = `rgba(139, 92, 246, ${alpha})`;
            ctx.lineWidth = 0.7;
            ctx.moveTo(p.x, p.y);
            ctx.lineTo(p2.x, p2.y);
            ctx.stroke();
          }
        }
      }

      rafRef.current = requestAnimationFrame(draw);
    };

    const handleMouseMove = (e: MouseEvent) => {
      const rect = canvas.getBoundingClientRect();
      mouseRef.current = { x: e.clientX - rect.left, y: e.clientY - rect.top };
    };
    const handleMouseLeave = () => {
      mouseRef.current = { x: -9999, y: -9999 };
    };

    resize();
    init();
    draw();

    const ro = new ResizeObserver(() => { resize(); init(); });
    ro.observe(canvas);
    canvas.addEventListener("mousemove", handleMouseMove);
    canvas.addEventListener("mouseleave", handleMouseLeave);

    return () => {
      cancelAnimationFrame(rafRef.current);
      ro.disconnect();
      canvas.removeEventListener("mousemove", handleMouseMove);
      canvas.removeEventListener("mouseleave", handleMouseLeave);
    };
  }, []);

  return (
    <canvas
      ref={canvasRef}
      className="absolute inset-0 w-full h-full"
      style={{ display: "block" }}
    />
  );
}
