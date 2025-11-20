import { Box, Paper, Typography, Alert } from "@mui/material";
import { useWeather } from "../hooks/useWeather";
import ThermostatIcon from "@mui/icons-material/Thermostat";

export default function WeatherResult() {
  const { result, error } = useWeather();

  // Não mostra nada se não houver resultado nem erro
  if (!result && !error) {
    return null;
  }

  // Mostra erro
  if (error) {
    return (
      <Paper elevation={2} sx={{ p: 3, bgcolor: "error.light" }}>
        <Alert severity="error">{error}</Alert>
      </Paper>
    );
  }

  // Mostra resultado
  if (result) {
    return (
      <Paper elevation={2} sx={{ p: 3 }}>
        <Box sx={{ display: "flex", alignItems: "center", gap: 1, mb: 2 }}>
          <ThermostatIcon color="primary" />
          <Typography variant="h5" component="h2">
            Temperatura
          </Typography>
        </Box>

        <Box
          sx={{
            display: "grid",
            gridTemplateColumns: "repeat(auto-fit, minmax(120px, 1fr))",
            gap: 2,
          }}
        >
          <Box>
            <Typography variant="body2" color="text.secondary">
              Celsius
            </Typography>
            <Typography variant="h4" color="primary.main" fontWeight="bold">
              {result.temp_C}°C
            </Typography>
          </Box>

          <Box>
            <Typography variant="body2" color="text.secondary">
              Fahrenheit
            </Typography>
            <Typography variant="h4" color="primary.main" fontWeight="bold">
              {result.temp_F}°F
            </Typography>
          </Box>

          <Box>
            <Typography variant="body2" color="text.secondary">
              Kelvin
            </Typography>
            <Typography variant="h4" color="primary.main" fontWeight="bold">
              {result.temp_K}K
            </Typography>
          </Box>
        </Box>
      </Paper>
    );
  }

  return null;
}
