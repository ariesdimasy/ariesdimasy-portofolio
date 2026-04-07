// Portfolio data - customize this with your actual information

export interface Education {
  institution: string;
  degree: string;
  field: string;
  startYear: string;
  endYear: string;
  description?: string;
}

export interface Experience {
  company: string;
  position: string;
  startDate: string;
  endDate: string;
  description: string;
  technologies: string[];
}

export interface Project {
  title: string;
  description: string;
  technologies: string[];
  imageUrl?: string;
  liveUrl?: string;
  githubUrl?: string;
}

export interface Skill {
  name: string;
  category: "Frontend" | "Backend" | "Tools" | "Database";
  icon?: string;
}

export interface Certificate {
  name: string;
  organization: string;
  year: string;
  credentialUrl?: string;
}

export const biodata = {
  name: "Aries Dimas Yudhistira",
  tagline: "Full Stack Web Developer",
  description:
    "Passionate developer with expertise in building modern web applications. I specialize in creating performant, scalable, and beautiful digital experiences using cutting-edge technologies.",
  email: "ariesdimasy@email.com",
  github: "https://github.com/ariesdimasy",
  linkedin: "https://linkedin.com/in/ariesdimasy",
};

export const educations: Education[] = [
  {
    institution: "Universitas Indonesia",
    degree: "Sarjana Komputer (S.Kom)",
    field: "Teknik Informatika",
    startYear: "2018",
    endYear: "2022",
    description:
      "Focused on software engineering and web development. Active in various campus technology communities and projects.",
  },
  {
    institution: "SMA Negeri 1 Jakarta",
    degree: "SMA",
    field: "IPA (Ilmu Pengetahuan Alam)",
    startYear: "2015",
    endYear: "2018",
    description:
      "Graduated with honors. Participated in science olympiad and computer programming club.",
  },
];

export const experiences: Experience[] = [
  {
    company: "PT Tech Solutions Indonesia",
    position: "Senior Full Stack Developer",
    startDate: "Jan 2024",
    endDate: "Sekarang",
    description:
      "Leading the development of enterprise web applications using React and Node.js. Mentoring junior developers and architecting scalable solutions for high-traffic systems.",
    technologies: ["React.js", "Node.js", "PostgreSQL", "Docker", "AWS"],
  },
  {
    company: "PT Digital Kreatif",
    position: "Full Stack Developer",
    startDate: "Jun 2022",
    endDate: "Des 2023",
    description:
      "Developed and maintained multiple client-facing web applications. Implemented RESTful APIs and optimized database queries for improved performance.",
    technologies: ["React.js", "PHP", "Laravel", "MySQL", "Redis"],
  },
  {
    company: "Startup XYZ",
    position: "Junior Web Developer",
    startDate: "Jan 2021",
    endDate: "Mei 2022",
    description:
      "Built responsive web interfaces and contributed to the development of an e-commerce platform serving thousands of daily users.",
    technologies: ["JavaScript", "Vue.js", "PHP", "MySQL"],
  },
];

export const projects: Project[] = [
  {
    title: "E-Commerce Platform",
    description:
      "A full-featured e-commerce platform with product management, shopping cart, payment integration, and real-time order tracking. Built with modern architecture for scalability.",
    technologies: ["React.js", "Node.js", "MongoDB", "Stripe", "Tailwind CSS"],
    liveUrl: "#",
    githubUrl: "#",
  },
  {
    title: "Project Management Dashboard",
    description:
      "A comprehensive project management tool with kanban boards, team collaboration, task assignment, and analytics dashboard with real-time updates.",
    technologies: ["React.js", "TypeScript", "PostgreSQL", "Socket.io"],
    liveUrl: "#",
    githubUrl: "#",
  },
  {
    title: "Blog CMS Platform",
    description:
      "A content management system for blogging with rich text editor, SEO optimization, user authentication, and multi-language support.",
    technologies: ["PHP", "Laravel", "MySQL", "Vue.js", "Tailwind CSS"],
    liveUrl: "#",
    githubUrl: "#",
  },
  {
    title: "Real-time Chat Application",
    description:
      "A modern real-time messaging application with group chats, file sharing, online status indicators, and end-to-end encryption.",
    technologies: ["React.js", "Node.js", "Socket.io", "Redis", "MongoDB"],
    liveUrl: "#",
    githubUrl: "#",
  },
];

export const skills: Skill[] = [
  { name: "React.js", category: "Frontend" },
  { name: "TypeScript", category: "Frontend" },
  { name: "JavaScript", category: "Frontend" },
  { name: "Vue.js", category: "Frontend" },
  { name: "Tailwind CSS", category: "Frontend" },
  { name: "HTML/CSS", category: "Frontend" },
  { name: "PHP", category: "Backend" },
  { name: "Laravel", category: "Backend" },
  { name: "Node.js", category: "Backend" },
  { name: "Express.js", category: "Backend" },
  { name: "Go", category: "Backend" },
  { name: "REST API", category: "Backend" },
  { name: "PostgreSQL", category: "Database" },
  { name: "MySQL", category: "Database" },
  { name: "MongoDB", category: "Database" },
  { name: "Redis", category: "Database" },
  { name: "Git", category: "Tools" },
  { name: "Docker", category: "Tools" },
  { name: "AWS", category: "Tools" },
  { name: "Linux", category: "Tools" },
];

export const certificates: Certificate[] = [
  {
    name: "AWS Certified Cloud Practitioner",
    organization: "Amazon Web Services",
    year: "2024",
    credentialUrl: "#",
  },
  {
    name: "Meta Front-End Developer Professional Certificate",
    organization: "Meta (via Coursera)",
    year: "2023",
    credentialUrl: "#",
  },
  {
    name: "Full Stack Web Development",
    organization: "Dicoding Indonesia",
    year: "2022",
    credentialUrl: "#",
  },
  {
    name: "Google IT Support Professional Certificate",
    organization: "Google (via Coursera)",
    year: "2021",
    credentialUrl: "#",
  },
];

export const navLinks = [
  { label: "Home", href: "#home" },
  { label: "Education", href: "#education" },
  { label: "Experience", href: "#experience" },
  { label: "Projects", href: "#projects" },
  { label: "Skills", href: "#skills" },
  { label: "Certificates", href: "#certificates" },
];
