"use client";

import { useEffect, useState } from "react";
import { getUser } from "@/lib/auth";
import { User } from "@/types";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import {
  User as UserIcon,
  Zap,
  GraduationCap,
  Briefcase,
  Award,
  FolderKanban,
  Share2,
  ArrowRight,
} from "lucide-react";
import Link from "next/link";

const menuCards = [
  {
    label: "Biodata",
    description: "Kelola info pribadi dan headline",
    href: "/dashboard/biodata",
    icon: UserIcon,
    color: "text-blue-400",
    bg: "bg-blue-500/10 border-blue-500/20",
  },
  {
    label: "Skills",
    description: "Tambah dan kelola keahlian teknis",
    href: "/dashboard/skills",
    icon: Zap,
    color: "text-yellow-400",
    bg: "bg-yellow-500/10 border-yellow-500/20",
  },
  {
    label: "Social Media",
    description: "Kelola link media sosial",
    href: "/dashboard/sosmeds",
    icon: Share2,
    color: "text-pink-400",
    bg: "bg-pink-500/10 border-pink-500/20",
  },
  {
    label: "Education",
    description: "Riwayat pendidikan formal",
    href: "/dashboard/education",
    icon: GraduationCap,
    color: "text-green-400",
    bg: "bg-green-500/10 border-green-500/20",
  },
  {
    label: "Experience",
    description: "Pengalaman kerja profesional",
    href: "/dashboard/experience",
    icon: Briefcase,
    color: "text-orange-400",
    bg: "bg-orange-500/10 border-orange-500/20",
  },
  {
    label: "Certificates",
    description: "Sertifikasi dan penghargaan",
    href: "/dashboard/certificates",
    icon: Award,
    color: "text-purple-400",
    bg: "bg-purple-500/10 border-purple-500/20",
  },
  {
    label: "Projects",
    description: "Portfolio project yang telah dibuat",
    href: "/dashboard/projects",
    icon: FolderKanban,
    color: "text-indigo-400",
    bg: "bg-indigo-500/10 border-indigo-500/20",
  },
];

export default function DashboardPage() {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    setUser(getUser());
  }, []);

  const hour = new Date().getHours();
  const greeting =
    hour < 12 ? "Selamat pagi" : hour < 17 ? "Selamat siang" : "Selamat malam";

  return (
    <div className="space-y-8">
      {/* Welcome banner */}
      <div className="relative rounded-2xl overflow-hidden bg-gradient-to-r from-indigo-600/20 via-indigo-600/10 to-transparent border border-indigo-500/20 p-6">
        <div className="absolute top-0 right-0 w-64 h-64 bg-indigo-500/5 rounded-full -translate-y-1/2 translate-x-1/2 pointer-events-none" />
        <div className="relative">
          <Badge className="bg-indigo-600/20 text-indigo-400 border border-indigo-500/30 mb-3">
            Admin Panel
          </Badge>
          <h2 className="text-2xl font-bold text-white">
            {greeting}, {user?.name ?? "Admin"}! 👋
          </h2>
          <p className="text-zinc-400 mt-1 text-sm">
            Kelola data portofolio kamu dari panel ini. Pilih menu di bawah untuk mulai.
          </p>
        </div>
      </div>

      {/* Menu grid */}
      <div>
        <h3 className="text-sm font-medium text-zinc-500 uppercase tracking-wider mb-4">
          Menu Manajemen
        </h3>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          {menuCards.map((item) => {
            const Icon = item.icon;
            return (
              <Link key={item.href} href={item.href}>
                <Card className="group h-full bg-zinc-900 border-zinc-800 hover:border-zinc-700 transition-all duration-200 cursor-pointer hover:shadow-lg hover:shadow-black/20 hover:-translate-y-0.5">
                  <CardHeader className="pb-3">
                    <div
                      className={`w-10 h-10 rounded-xl border flex items-center justify-center mb-3 ${item.bg} transition-transform group-hover:scale-110`}
                    >
                      <Icon className={`w-5 h-5 ${item.color}`} />
                    </div>
                    <CardTitle className="text-base font-semibold text-white flex items-center justify-between">
                      {item.label}
                      <ArrowRight className="w-4 h-4 text-zinc-600 group-hover:text-zinc-400 group-hover:translate-x-1 transition-all" />
                    </CardTitle>
                  </CardHeader>
                  <CardContent className="pt-0">
                    <p className="text-xs text-zinc-500">{item.description}</p>
                  </CardContent>
                </Card>
              </Link>
            );
          })}
        </div>
      </div>
    </div>
  );
}
