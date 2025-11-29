import { Box, Container, Paper, Typography } from "@mui/material";
import WeatherForm from "./components/WeatherForm";
import WeatherResult from "./components/WeatherResult";
import CloudIcon from "@mui/icons-material/Cloud";
import ThemeToggle from "./components/ThemeToggle";
import { useWeather } from "./hooks/useWeather";
import { dayTheme, nightTheme } from "./theme/theme";

function App() {
  const { theme, toggleTheme } = useWeather();
  const isDay = theme === 'day';
  const currentTheme = isDay ? dayTheme : nightTheme;

  return (
    <Box
      className={`animated-background animated-background--${theme}`}
      sx={{
        minHeight: '100vh',
        width: '100vw',
        display: 'flex',
        alignItems: 'flex-start',
        justifyContent: 'center',
        p: { xs: 2, sm: 3, md: 4 },
        pt: { xs: 3, sm: 4, md: 6 },
        backgroundImage: currentTheme.background.gradient,
      }}
    >
        <Container maxWidth="sm" sx={{ p: 0, position: 'relative' }}>
          <Box sx={{
            display: 'flex',
            flexDirection: 'column',
            gap: 3,
            alignItems: 'center'
          }}>
            {/* App Title */}
            <Box sx={{ textAlign: 'center', width: '100%' }}>
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
                  mb: 2.5
                }}
              >
                Informe o CEP para consultar o clima
              </Typography>

              {/* Theme Toggle - Centered below subtitle */}
              <Box sx={{
                display: 'flex',
                justifyContent: 'center',
                mt: 1
              }}>
                <ThemeToggle
                  checked={!isDay}
                  onChange={toggleTheme}
                  scale={1.5}
                />
              </Box>
            </Box>

            {/* Search Form */}
            <Paper
              elevation={0}
              className={`glass-panel glass-panel--${theme} transition-colors`}
              sx={{
                width: '100%',
                p: { xs: 3, sm: 4 },
                borderRadius: 4,
                backgroundColor: 'transparent',
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
