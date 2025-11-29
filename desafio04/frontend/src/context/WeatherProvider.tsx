import { useState, useCallback, useMemo } from "react";
import type { ReactNode } from "react";
import WeatherContext, { type WeatherData, type Theme } from "./WeatherContext";
import { fetchWeatherByCep } from "../services/weatherService";

interface WeatherContextProviderProps {
  children: ReactNode;
}

export default function WeatherProvider(
  { children }: WeatherContextProviderProps
) {
  const [error, setError] = useState<string | null>(null);
  const [result, setResult] = useState<WeatherData | null>(null);
  const [theme, setTheme] = useState<Theme>('day');

  const searchWeather = useCallback(async (cep: string) => {
    // Limpar estados anteriores
    setError(null);
    setResult(null);

    try {
      const data = await fetchWeatherByCep(cep);
      setResult(data);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Erro ao buscar clima';
      setError(errorMessage);
    }
  }, []);

  const toggleTheme = useCallback(() => {
    setTheme(prev => prev === 'day' ? 'night' : 'day');
  }, []);

  const value = useMemo(
    () => ({
      error,
      result,
      searchWeather,
      theme,
      toggleTheme,
    }),
    [error, result, searchWeather, theme, toggleTheme]
  );

  return (
    <WeatherContext.Provider value={value}>
      {children}
    </WeatherContext.Provider>
  );
}
