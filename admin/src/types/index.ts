// Types untuk semua entitas portfolio

export interface User {
  id: number;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface Biodata {
  id: number;
  user_id: number;
  headline: string;
  about: string;
  address: string;
  phone: string;
  created_at: string;
  updated_at: string;
}

export interface Skill {
  id: number;
  name: string;
  icon: string;
  type: "frontend" | "backend" | "database";
  created_at: string;
  updated_at: string;
}

export interface Sosmed {
  id: number;
  user_id: number;
  name: string;
  icon: string;
  link: string;
  created_at: string;
  updated_at: string;
}

export interface Education {
  id: number;
  user_id: number;
  degree: string;
  major: string;
  institution: string;
  start_date: string;
  end_date: string;
  created_at: string;
  updated_at: string;
}

export interface Experience {
  id: number;
  user_id: number;
  title: string;
  company_name: string;
  start_date: string;
  end_date: string;
  description: string;
  skill: Skill[];
  created_at: string;
  updated_at: string;
}

export interface Certificate {
  id: number;
  user_id: number;
  name: string;
  organization: string;
  credential_id: string;
  credential_url: string;
  issue_date: string;
  expiration_date: string;
  image: string;
  created_at: string;
  updated_at: string;
}

export interface Project {
  id: number;
  user_id: number;
  title: string;
  description: string;
  image: string;
  link: string;
  skill: Skill[];
  project_image: ProjectImage[];
  created_at: string;
  updated_at: string;
}

export interface ProjectImage {
  id: number;
  project_id: number;
  image: string;
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface LoginResponse {
  message: string;
  token: string;
  user: User;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
}
