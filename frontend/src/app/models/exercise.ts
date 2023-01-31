export interface Exercise {
    id: number;
    name?: String;
    description?: String;
    img?: String;
    coachId?: number
  }
  
export interface CreateExerciseDTO {
    name: String;
    description: String;
    img: String;
    coachId?: number
}