import { createTheme } from '@mui/material/styles';

// Day Theme - Clear morning sky
export const dayTheme = {
  name: 'day',
  background: {
    gradient: 'linear-gradient(-45deg, #4facfe, #00f2fe, #4facfe, #00d4ff)',
  },
  text: {
    primary: '#1a1a2e',
    secondary: 'rgba(26, 26, 46, 0.8)',
    onBackground: 'rgba(255, 255, 255, 0.98)',
    onGlass: '#1a1a2e',
  },
  button: {
    gradient: 'linear-gradient(135deg, #667eea, #764ba2, #667eea, #764ba2)',
    text: '#ffffff',
  },
  icon: '#667eea',
  switch: {
    background: 'rgba(26, 26, 46, 0.3)',
    emoji: '#1a1a2e',
  },
  input: {
    background: 'rgba(255, 255, 255, 0.25)',
    border: 'rgba(26, 26, 46, 0.3)',
    borderHover: 'rgba(26, 26, 46, 0.5)',
    borderFocus: 'rgba(26, 26, 46, 0.7)',
  },
  divider: 'rgba(26, 26, 46, 0.25)',
  glass: {
    border: 'rgba(26, 26, 46, 0.25)',
  },
  glassForm: {
    border: 'rgba(26, 26, 46, 0.25)',
  },
  glassResult: {
    border: 'rgba(26, 26, 46, 0.25)',
  },
  glassCard: {
    border: 'rgba(26, 26, 46, 0.3)',
  },
};

// Night Theme - Deep twilight
export const nightTheme = {
  name: 'night',
  background: {
    gradient: 'linear-gradient(-45deg, #141e30, #243b55, #141e30, #0f2027)',
  },
  text: {
    primary: '#ffffff',
    secondary: 'rgba(255, 255, 255, 0.9)',
    onBackground: 'rgba(255, 255, 255, 0.98)',
    onGlass: '#ffffff',
  },
  button: {
    gradient: 'linear-gradient(135deg, #f093fb, #f5576c, #f093fb, #f5576c)',
    text: '#ffffff',
  },
  icon: '#f093fb',
  switch: {
    background: 'rgba(255, 255, 255, 0.15)',
    emoji: '#ffffff',
  },
  input: {
    background: 'rgba(255, 255, 255, 0.25)',
    border: 'rgba(255, 255, 255, 0.4)',
    borderHover: 'rgba(255, 255, 255, 0.6)',
    borderFocus: 'rgba(255, 255, 255, 0.8)',
  },
  divider: 'rgba(255, 255, 255, 0.3)',
  glass: {
    border: 'rgba(255, 255, 255, 0.35)',
  },
  glassForm: {
    border: 'rgba(255, 255, 255, 0.35)',
  },
  glassResult: {
    border: 'rgba(255, 255, 255, 0.35)',
  },
  glassCard: {
    border: 'rgba(255, 255, 255, 0.40)',
  },
};

// MUI Base Theme
const theme = createTheme({
  palette: {
    primary: {
      main: '#5e35b1',
      contrastText: '#fff',
    },
    secondary: {
      main: '#2F80ED',
      contrastText: '#fff',
    },
  },
  typography: {
    fontFamily: '"Inter", "Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      fontWeight: 700,
    },
    h2: {
      fontWeight: 700,
    },
    h4: {
      fontWeight: 600,
    },
  },
  shape: {
    borderRadius: 12,
  },
});

export default theme;
