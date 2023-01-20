export interface Exercise {
    id: number;
    name?: String;
    description?: String;
    img?: String;
  }
  
export interface CreateExerciseDTO {
    name: String;
    description: String;
    img: String;
}