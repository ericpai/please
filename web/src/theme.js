import { red, green } from "@material-ui/core/colors";
import { createMuiTheme } from "@material-ui/core/styles";
import { zhCN } from "@material-ui/core/locale";

// A custom theme for this app
const theme = createMuiTheme({
  palette: {
    primary: {
      main: "#556cd6",
    },
    secondary: {
      main: "#19857b",
    },
    error: {
      main: red.A400,
    },
    background: {
      default: "#fff",
    },
    new: {
      main: green[700],
    },
    edit: {
      main: green[700],
    },
    delete: {
      main: red[700],
    },
  },
}, zhCN);

export default theme;
