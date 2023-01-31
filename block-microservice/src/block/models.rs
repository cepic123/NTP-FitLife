use serde::Deserialize;

#[derive(Deserialize, Clone)]
pub struct CreateBlockDTO {
    pub user_id: i32,
    pub user_name: String,
    pub subject_name: String,
    pub block_subject_id: i32,
}
