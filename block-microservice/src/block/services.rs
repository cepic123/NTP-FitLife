use actix_web::{get, post, web::{Data, Json, Path, ServiceConfig}, Responder, HttpResponse};
use crate::{AppState, Block};
use super::models::{CreateBlockDTO};
use sqlx::{self};

#[get("/block/{id}")]
async fn get_block(state: Data<AppState>, path: Path<i32>) -> impl Responder {
    
    let id: i32 = path.into_inner();

    match sqlx::query_as::<_,Block>(
        "SELECT id, user_id, user_name, subject_name, block_subject_id FROM blocks WHERE id = $1"
    )
        .bind(id)
        .fetch_one(&state.db)
        .await
    {
        Ok(blocks) => HttpResponse::Ok().json(blocks),
        Err(_) => HttpResponse::NotFound().json("No block found"),
    }
}

#[get("/block")]
async fn get_blocks(state: Data<AppState>) -> impl Responder {
    println!("hello there!");

    match sqlx::query_as::<_,Block>(
        "SELECT id, user_id, user_name, subject_name, block_subject_id FROM blocks"
    )
        .fetch_all(&state.db)
        .await
    {
        Ok(blocks) => HttpResponse::Ok().json(blocks),
        Err(_) => HttpResponse::NotFound().json("No blocks found"),
    }
}

#[get("/block/user/{id}")]
async fn get_user_blocks(state: Data<AppState>, path: Path<i32>) -> impl Responder {
    
    let id: i32 = path.into_inner();

    match sqlx::query_as::<_,Block>(
        "SELECT id, user_id, user_name, subject_name, block_subject_id FROM blocks WHERE user_id = $1"
    )
        .bind(id)
        .fetch_all(&state.db)
        .await
    {
        Ok(blocks) => HttpResponse::Ok().json(blocks),
        Err(_) => HttpResponse::NotFound().json("No blocks found by user"),
    }
}

#[get("/block/subject/{id}")]
async fn get_subject_blocks(state: Data<AppState>, path: Path<i32>) -> impl Responder {
    
    let id: i32 = path.into_inner();

    match sqlx::query_as::<_,Block>(
        "SELECT id, user_id, user_name, subject_name, block_subject_id FROM blocks WHERE block_subject_id = $1"
    )
        .bind(id)
        .fetch_all(&state.db)
        .await
    {
        Ok(blocks) => HttpResponse::Ok().json(blocks),
        Err(_) => HttpResponse::NotFound().json("No blocks found for subject"),
    }
}

#[post("/block")]
async fn create_block(state: Data<AppState>, body: Json<CreateBlockDTO>) -> impl Responder {
    match sqlx::query_as::<_, Block>(
        "INSERT INTO blocks (user_id, user_name, subject_name, block_subject_id) VALUES ($1, $2, $3, $4) RETURNING id, user_id, user_name, subject_name, block_subject_id"
    )
        .bind(body.user_id)
        .bind(body.user_name.to_string())
        .bind(body.subject_name.to_string())
        .bind(body.block_subject_id)
        .fetch_one(&state.db)
        .await
        {
            Ok(block) => HttpResponse::Ok().json(block),
            Err(_) => HttpResponse::InternalServerError().json("Failed to create block"),
        }
}

pub fn config(cfg: &mut ServiceConfig) {
    cfg.service(get_block)
    .service(get_blocks)
    .service(create_block)
    .service(get_subject_blocks)
    .service(get_user_blocks);
}
