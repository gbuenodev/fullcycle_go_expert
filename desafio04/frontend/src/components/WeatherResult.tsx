import { Box, Paper, Typography, Alert, Divider, Card, CardContent } from "@mui/material";
import { useWeather } from "../hooks/useWeather";
import PlaceIcon from "@mui/icons-material/Place";
import ThermostatIcon from "@mui/icons-material/Thermostat";
import { dayTheme, nightTheme } from "../theme/theme";


export default function WeatherResult() {
  const { result, error, theme } = useWeather();
  const isDay = theme === 'day';
  const currentTheme = isDay ? dayTheme : nightTheme;

  // Não mostra nada se não houver resultado nem erro
  if (!result && !error) {
    return null;
  }

  // Mostra erro
  if (error) {
    return (
      <Paper
        elevation={0}
        className="glass-result"
        sx={{
          p: 3,
          borderRadius: 4,
          background: 'transparent',
          border: `1px solid ${currentTheme.glassResult.border}`,
        }}
      >
        <Alert
          severity="error"
          className="text-readable"
          sx={{
            background: 'rgba(211, 47, 47, 0.85)',
            backdropFilter: 'blur(15px)',
            WebkitBackdropFilter: 'blur(15px)',
            border: '1px solid rgba(255, 255, 255, 0.3)',
            color: 'white',
            fontWeight: 'bold',
            borderRadius: 2,
          }}
        >
          {error}
        </Alert>
      </Paper>
    );
  }

  // Mostra resultado
  if (result) {
    return (
      <Paper
        elevation={0}
        className="glass-result transition-colors"
        sx={{
          p: { xs: 3, sm: 4 },
          borderRadius: 4,
          color: currentTheme.text.onGlass,
          background: 'transparent',
          border: `1px solid ${currentTheme.glassResult.border}`,
        }}
      >
        {/* Location Header */}
        <Box sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          gap: 1.5,
          mb: 3
        }}>
          <PlaceIcon
            className="transition-colors"
            sx={{
              fontSize: { xs: 28, sm: 32 },
              color: currentTheme.icon,
            }}
          />
          <Typography
            variant="h4"
            component="h2"
            fontWeight="bold"
            className="transition-colors text-readable"
            sx={{
              fontSize: { xs: '1.5rem', sm: '2rem' },
              color: currentTheme.text.onGlass,
            }}
          >
            {result.city}
          </Typography>
        </Box>

        <Divider
          className="transition-colors"
          sx={{
            borderColor: currentTheme.divider,
            mb: 3,
          }}
        />

        {/* Primary Temperature - Celsius */}
        <Box sx={{ textAlign: 'center', mb: 3 }}>
          <Box sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            gap: 1,
            mb: 1.5
          }}>
            <ThermostatIcon
              className="transition-colors"
              sx={{
                fontSize: { xs: 24, sm: 28 },
                color: currentTheme.icon,
              }}
            />
            <Typography
              variant="h6"
              className="transition-colors text-readable"
              sx={{
                color: currentTheme.text.onGlass,
                fontSize: { xs: '1rem', sm: '1.25rem' },
                fontWeight: 600,
                opacity: 0.9,
              }}
            >
              Temperatura Atual
            </Typography>
          </Box>
          <Typography
            variant="h1"
            className="transition-colors text-readable"
            sx={{
              fontSize: { xs: '3.5rem', sm: '4.5rem', md: '5.5rem' },
              fontWeight: 'bold',
              lineHeight: 1,
              color: currentTheme.text.onGlass,
              mb: 1,
            }}
          >
            {result.temp_C}°
          </Typography>
          <Typography
            variant="h5"
            className="transition-colors text-readable"
            sx={{
              color: currentTheme.text.onGlass,
              fontSize: { xs: '1.25rem', sm: '1.5rem' },
              fontWeight: 500,
              opacity: 0.9,
            }}
          >
            Celsius
          </Typography>
        </Box>

        {/* Secondary Temperatures - Cards */}
        <Box
          sx={{
            display: 'grid',
            gridTemplateColumns: { xs: '1fr', sm: 'repeat(2, 1fr)' },
            gap: 2
          }}
        >
          <Card
            elevation={0}
            className="glass-card"
            sx={{
              borderRadius: 3,
              background: 'transparent',
              border: `1px solid ${currentTheme.glassCard.border}`,
            }}
          >
            <CardContent sx={{ textAlign: 'center', py: 3, px: 2 }}>
              <Typography
                variant="body2"
                className="transition-colors text-readable"
                sx={{
                  color: currentTheme.text.onGlass,
                  opacity: 0.9,
                }}
                gutterBottom
                fontWeight="medium"
              >
                Fahrenheit
              </Typography>
              <Typography
                variant="h3"
                className="transition-colors text-readable"
                sx={{
                  color: currentTheme.text.onGlass,
                }}
                fontWeight="bold"
              >
                {result.temp_F}°
              </Typography>
            </CardContent>
          </Card>

          <Card
            elevation={0}
            className="glass-card"
            sx={{
              borderRadius: 3,
              background: 'transparent',
              border: `1px solid ${currentTheme.glassCard.border}`,
            }}
          >
            <CardContent sx={{ textAlign: 'center', py: 3, px: 2 }}>
              <Typography
                variant="body2"
                className="transition-colors text-readable"
                sx={{
                  color: currentTheme.text.onGlass,
                  opacity: 0.9,
                }}
                gutterBottom
                fontWeight="medium"
              >
                Kelvin
              </Typography>
              <Typography
                variant="h3"
                className="transition-colors text-readable"
                sx={{
                  color: currentTheme.text.onGlass,
                }}
                fontWeight="bold"
              >
                {result.temp_K}K
              </Typography>
            </CardContent>
          </Card>
        </Box>
      </Paper>
    );
  }

  return null;
}
