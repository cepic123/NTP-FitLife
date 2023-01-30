export interface Role {
    name: string,
    code: string
}

export interface User {
    id: number,
    email: string,
    username: string,
    role: string,
    DeletedAt: string,
    CreatedAt: string
}