import { useState, useCallback, useMemo } from "react";
import type { ReactNode } from "react";
import WeatherContext, { type WeatherData } from "./WeatherContext";
import { fetchWeatherByCep } from "../services/weatherService";

interface WeatherContextProviderProps {
  children: ReactNode;
}

export default function WeatherProvider(
  { children }: WeatherContextProviderProps
) {
  const [error, setError] = useState<string | null>(null);
  const [result, setResult] = useState<WeatherData | null>(null);

  const searchWeather = useCallback(async (cep: string) => {
    console.log('Buscando clima para o CEP:', cep);

    // Limpar estados anteriores
    setError(null);
    setResult(null);

    try {
      const data = await fetchWeatherByCep(cep);
      setResult(data);
      console.log('Clima recebido:', data);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Erro ao buscar clima';
      setError(errorMessage);
      console.error('Erro ao buscar clima:', errorMessage);
    }
  }, []);

  const value = useMemo(
    () => ({
      error,
      result,
      searchWeather,
    }),
    [error, result, searchWeather]
  );

  return (
    <WeatherContext.Provider value={value}>
      {children}
    </WeatherContext.Provider>
  );
}
