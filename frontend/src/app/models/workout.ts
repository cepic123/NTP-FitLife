import { Exercise } from "./exercise"

export interface Workout {
    ID?: number,
    name?: string,
    description?: string
    sets: Set[]
    rating?: number
    coachId?: number
}

export interface Set {
    orderNum: number,
    noSets?: number,
    breakLngth?: number,
    reps: Rep[]
}

export interface Rep {
    orderNum: number,
    noReps: number,
    exercise?: Exercise
}