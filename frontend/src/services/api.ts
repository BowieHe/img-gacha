import axios from 'axios'
import type { GenerationRequest, GenerationResponse, GenerationStatus } from '@/types'

const API_BASE_URL = '/api'

const api = axios.create({
    baseURL: API_BASE_URL,
    timeout: 60000,
})

export const apiService = {
    generateImages: async (request: GenerationRequest): Promise<GenerationResponse> => {
        const response = await api.post('/generate', request)
        return response.data
    },

    getGenerationStatus: async (taskId: string): Promise<GenerationStatus> => {
        const response = await api.get(`/status/${taskId}`)
        return response.data
    },

    getAvailableModels: async (): Promise<string[]> => {
        const response = await api.get('/models')
        return response.data.models
    },

    health: async (): Promise<{ status: string }> => {
        const response = await api.get('/health')
        return response.data
    },
}

export default apiService
