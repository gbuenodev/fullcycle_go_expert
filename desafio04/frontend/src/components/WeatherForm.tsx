import { useState, type FormEvent } from "react";
import { Box, TextField, Button } from "@mui/material";
import InputMask from "react-input-mask";
import { useWeather } from "../hooks/useWeather";
import { validateCep } from "../utils/validation";
import Loading from "./Loading";
import { dayTheme, nightTheme } from "../theme/theme";

export default function WeatherForm() {
  const [cep, setCep] = useState<string>("");
  const [cepError, setCepError] = useState<string>("");
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const { searchWeather, theme } = useWeather();
  const isDay = theme === 'day';
  const currentTheme = isDay ? dayTheme : nightTheme;

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
      <Box sx={{ display: "flex", flexDirection: "column", gap: 2.5 }}>
        <InputMask
          mask="99999-999"
          value={cep}
          onChange={handleCepChange}
          disabled={isSubmitting}
        >
          {(inputProps: any) => (
            <TextField
              {...inputProps}
              label="CEP"
              placeholder="00000-000"
              fullWidth
              variant="outlined"
              error={!!cepError}
              helperText={cepError || "Digite o CEP"}
              sx={{
                '& .MuiOutlinedInput-root': {
                  fontSize: '1.1rem',
                  background: currentTheme.input.background,
                  backdropFilter: 'blur(15px)',
                  WebkitBackdropFilter: 'blur(15px)',
                  borderRadius: 3,
                  '& input': {
                    color: currentTheme.text.onGlass,
                    fontWeight: 500,
                    '&::placeholder': {
                      color: currentTheme.text.secondary,
                      opacity: 0.7,
                    }
                  },
                  '& fieldset': {
                    borderColor: currentTheme.input.border,
                    borderWidth: 1.5,
                  },
                  '&:hover fieldset': {
                    borderColor: currentTheme.input.borderHover,
                  },
                  '&.Mui-focused fieldset': {
                    borderColor: currentTheme.input.borderFocus,
                    borderWidth: 2
                  }
                },
                '& .MuiInputLabel-root': {
                  color: currentTheme.text.onGlass,
                  fontWeight: 600,
                  '&.Mui-focused': {
                    color: currentTheme.text.onGlass,
                    fontWeight: 700
                  }
                },
                '& .MuiFormHelperText-root': {
                  color: currentTheme.text.secondary,
                  fontWeight: 500,
                }
              }}
              className="transition-colors"
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
            className="animated-button"
            sx={{
              py: 1.8,
              fontSize: '1.1rem',
              fontWeight: 'bold',
              borderRadius: 2,
              color: currentTheme.button.text,
              boxShadow: '0 4px 16px rgba(31, 38, 135, 0.3)',
              background: currentTheme.button.gradient,
              '&:hover': {
                transform: 'translateY(-2px)',
                boxShadow: '0 8px 24px rgba(0, 0, 0, 0.3)',
                filter: 'brightness(1.1)'
              },
              '&:disabled': {
                background: 'rgba(200, 200, 200, 0.5) !important',
                color: 'rgba(255, 255, 255, 0.5)',
                animation: 'none !important'
              },
              transition: 'transform 0.3s ease, box-shadow 0.3s ease, filter 0.3s ease'
            }}
          >
            Buscar Clima
          </Button>
        </Loading>
      </Box>
    </form>
  );
}
