import { combineReducers } from "redux";
import alertReducer, { AlertState } from "./alert";

const reducer = combineReducers({
  alert: alertReducer,
});

export interface AppState {
  alert: AlertState;
}

export default reducer;
