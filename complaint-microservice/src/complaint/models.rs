use serde::Deserialize;

#[derive(Deserialize, Clone)]
pub struct CreateComplaintDTO {
    pub user_id: usize,
    pub complaint_subject_id: usize,
    pub complaint_text: String,
}

#[derive(Deserialize, Clone)]
pub struct UpdateComplaintDTO {
    pub complaint_text: String,
}