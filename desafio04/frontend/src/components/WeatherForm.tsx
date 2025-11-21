import { useState, type FormEvent } from "react";
import { Box, TextField, Button } from "@mui/material";
import InputMask from "react-input-mask";
import { useWeather } from "../hooks/useWeather";
import { validateCep } from "../utils/validation";
import Loading from "./Loading";

export default function WeatherForm() {
  const [cep, setCep] = useState<string>("");
  const [cepError, setCepError] = useState<string>("");
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const { searchWeather } = useWeather();

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // Validar CEP com Zod
    const validation = validateCep(cep);

    if (!validation.success) {
      setCepError(validation.error || "CEP inválido");
      return;
    }

    // Limpar erro se passou na validação
    setCepError("");

    setIsSubmitting(true);
    try {
      // Envia o CEP limpo (sem hífen) para a API
      await searchWeather(validation.data!);
    } finally {
      setIsSubmitting(false);
    }
  };

  // Limpar erro quando usuário digita
  const handleCepChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCep(e.target.value);
    if (cepError) {
      setCepError("");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
        <InputMask
          mask="99999-999"
          value={cep}
          onChange={handleCepChange}
          disabled={isSubmitting}
        >
          {/* @ts-ignore - InputMask types issue */}
          {() => (
            <TextField
              label="CEP"
              placeholder="00000-000"
              fullWidth
              variant="outlined"
              error={!!cepError}
              helperText={cepError || "Digite o CEP (somente números)"}
              disabled={isSubmitting}
            />
          )}
        </InputMask>

        <Loading loading={isSubmitting}>
          <Button
            type="submit"
            variant="contained"
            fullWidth
            size="large"
            disabled={isSubmitting || !cep.trim()}
          >
            Buscar Clima
          </Button>
        </Loading>
      </Box>
    </form>
  );
}
