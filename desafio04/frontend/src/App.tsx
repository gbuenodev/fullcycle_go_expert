import { Box, Container, Paper, Typography } from "@mui/material";
import WeatherProvider from "./context/WeatherProvider";
import WeatherForm from "./components/WeatherForm";
import WeatherResult from "./components/WeatherResult";

function App() {
  return (
    <WeatherProvider>
      <Container maxWidth="sm">
        <Box sx={{
          minHeight: '100vh',
          backgroundColor: 'background.default',
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          gap: 3,
          py: 4
        }}>
          <Typography
            variant="h3" 
            component="h1"
            gutterBottom
            sx={{
            fontWeight: 'bold',
            color: 'primary.main'
          }}> 
            Consulta Clima ☀️
          </Typography>
          
          <Paper
            elevation={3}
            sx={{
              width: '100%',
              p: 4
            }}
          >
            <WeatherForm />
          </Paper>

          <Box sx={{ width: '100%' }}>
            <WeatherResult />
          </Box>
        </Box>
      </Container>
    </WeatherProvider>
  );
}

export default App;
