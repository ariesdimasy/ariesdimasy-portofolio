"use client";

import { usePathname } from "next/navigation";

const routeLabels: Record<string, string> = {
  "/dashboard": "Dashboard",
  "/dashboard/biodata": "Biodata",
  "/dashboard/skills": "Skills",
  "/dashboard/sosmeds": "Social Media",
  "/dashboard/education": "Education",
  "/dashboard/experience": "Experience",
  "/dashboard/certificates": "Certificates",
  "/dashboard/projects": "Projects",
};

export default function Header() {
  const pathname = usePathname();
  const title = routeLabels[pathname] ?? "Dashboard";

  return (
    <header className="h-16 flex items-center justify-between px-6 border-b border-zinc-800 bg-zinc-950/50 backdrop-blur-sm">
      <div>
        <h1 className="text-lg font-semibold text-white">{title}</h1>
        <p className="text-xs text-zinc-500">Kelola data portofolio kamu</p>
      </div>
    </header>
  );
}
