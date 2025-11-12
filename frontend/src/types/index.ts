export interface GenerationRequest {
    prompt: string
    model: string
    count: number
    seed?: number
    steps?: number
    cfg_scale?: number
    negative_prompt?: string
    height?: number
    width?: number
}

export interface GenerationResponse {
    task_id: string
    status: string
    created_at: string
}

export interface ImageResult {
    id: string
    task_id: string
    url: string
    seed: number
    created_at: string
}

export type AIModel = 'stable-diffusion-v1.5' | 'stable-diffusion-v2.1' | 'dall-e-3'

export interface GenerationStatus {
    task_id: string
    status: 'pending' | 'processing' | 'completed' | 'failed'
    results?: ImageResult[]
    error?: string
}
