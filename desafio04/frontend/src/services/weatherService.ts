import api from '../api';
import { AxiosError } from 'axios';

export interface WeatherResponse {
  city: string; 
  temp_C: number;
  temp_F: number;
  temp_K: number;
}

export async function fetchWeatherByCep(cep: string): Promise<WeatherResponse> {
  try {
    const { data } = await api.post<WeatherResponse>('/weather', { cep });
    return data;
  } catch (error) {
    if (error instanceof AxiosError) {
      // HTTP 422 - CEP inválido
      if (error.response?.status === 422) {
        throw new Error('CEP inválido');
      }

      // HTTP 404 - CEP não encontrado
      if (error.response?.status === 404) {
        throw new Error('CEP não encontrado');
      }

      // Outros erros HTTP
      if (error.response) {
        throw new Error(error.response.data?.message || 'Erro ao buscar dados do clima');
      }

      // Erro de rede ou timeout
      if (error.request) {
        throw new Error('Não foi possível conectar ao servidor');
      }
    }

    // Erro desconhecido
    throw new Error('Internal Server Error');
  }
}
