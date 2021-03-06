export interface Task {
  id: string;
  address: string;
  user: string;
  password: string;
  sourcePath: string;
  destPath: string;
  backend: string;
  schedule: string;
  succeed?: boolean;
  enabled: boolean;
  createdTime?: number;
  updatedTime?: number;
}

export interface Request {
  task: Task;
}

export interface Meta {
  code: number;
  message: string;
}

export interface Data {
  tasks: Task[];
}

export interface Response {
  meta: Meta;
  data: Data;
}
