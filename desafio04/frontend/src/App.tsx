import { Box, Container, Paper, Typography, Switch } from "@mui/material";
import WeatherForm from "./components/WeatherForm";
import WeatherResult from "./components/WeatherResult";
import CloudIcon from "@mui/icons-material/Cloud";
import { useWeather } from "./hooks/useWeather";
import { dayTheme, nightTheme } from "./theme/theme";

function App() {
  const { theme, toggleTheme } = useWeather();
  const isDay = theme === 'day';
  const currentTheme = isDay ? dayTheme : nightTheme;

  return (
    <Box
      className="animated-background"
      sx={{
        minHeight: '100vh',
        width: '100vw',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        p: { xs: 2, sm: 3, md: 4 },
        background: currentTheme.background.gradient,
      }}
    >
        <Container maxWidth="sm" sx={{ p: 0, position: 'relative' }}>
          {/* Theme Toggle - Fixed Position */}
          <Box
            className="glass transition-colors"
            sx={{
              position: 'fixed',
              top: { xs: 16, sm: 24 },
              right: { xs: 16, sm: 24 },
              display: 'flex',
              alignItems: 'center',
              gap: 1,
              p: 1.5,
              borderRadius: 3,
              zIndex: 1000,
              background: currentTheme.switch.background,
              border: `1px solid ${currentTheme.glass.border}`,
            }}
          >
            <Box
              className="transition-colors"
              sx={{
                fontSize: '1.25rem',
                color: currentTheme.switch.emoji,
              }}
            >
              ☀️
            </Box>
            <Switch
              checked={!isDay}
              onChange={toggleTheme}
              sx={{
                '& .MuiSwitch-switchBase': {
                  color: currentTheme.icon,
                  '&.Mui-checked': {
                    color: currentTheme.icon,
                  },
                  '&.Mui-checked + .MuiSwitch-track': {
                    backgroundColor: 'rgba(255, 255, 255, 0.4)',
                  },
                },
                '& .MuiSwitch-track': {
                  backgroundColor: 'rgba(255, 255, 255, 0.3)',
                },
              }}
            />
            <Box
              className="transition-colors"
              sx={{
                fontSize: '1.25rem',
                color: currentTheme.switch.emoji,
              }}
            >
              🌙
            </Box>
          </Box>

          <Box sx={{
            display: 'flex',
            flexDirection: 'column',
            gap: 3,
            alignItems: 'center'
          }}>
            {/* App Title */}
            <Box sx={{ textAlign: 'center', width: '100%', mt: 2 }}>

              <Box sx={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                gap: 2,
                mb: 1.5
              }}>
                <CloudIcon
                  className="text-glow"
                  sx={{
                    fontSize: { xs: 40, sm: 48 },
                    color: currentTheme.text.onBackground
                  }}
                />
                <Typography
                  variant="h2"
                  component="h1"
                  className="text-glow transition-colors"
                  sx={{
                    fontSize: { xs: '2rem', sm: '2.5rem', md: '3rem' },
                    fontWeight: 'bold',
                    color: currentTheme.text.onBackground,
                  }}
                >
                  Consulta Clima
                </Typography>
              </Box>
              <Typography
                variant="h6"
                className="text-shadow transition-colors"
                sx={{
                  fontSize: { xs: '1rem', sm: '1.25rem' },
                  color: currentTheme.text.onBackground,
                }}
              >
                Informe o CEP para consultar o clima
              </Typography>
            </Box>

            {/* Search Form */}
            <Paper
              elevation={0}
              className="glass-form"
              sx={{
                width: '100%',
                p: { xs: 3, sm: 4 },
                borderRadius: 4,
                background: 'transparent',
                border: `1px solid ${currentTheme.glassForm.border}`,
              }}
            >
              <WeatherForm />
            </Paper>

            {/* Results */}
            <Box sx={{ width: '100%' }}>
              <WeatherResult />
            </Box>
          </Box>
        </Container>
      </Box>
  );
}

export default App;
