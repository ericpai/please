import { combineReducers } from "redux";
import alertReducer, { AlertState } from "./alert";
import singleTaskReducer, { SingleTaskState } from "./task";

const reducer = combineReducers({
  alert: alertReducer,
  singleTask: singleTaskReducer,
});

export interface AppState {
  alert: AlertState;
  singleTask: SingleTaskState;
}

export default reducer;
