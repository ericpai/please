export interface AlertState {
  title: string;
  content: string;
  onConfirm: any;
  open: boolean;
}

export const OPEN_ALERT = "OPEN_ALERT";
export const CLOSE_ALERT = "CLOSE_ALERT";

export function openAlert(alert: AlertState): AlertActions {
  return {
    type: OPEN_ALERT,
    alert: alert,
  };
}

export function closeAlert(): AlertActions {
  return {
    type: CLOSE_ALERT,
  };
}

interface OpenAlertAction {
  type: typeof OPEN_ALERT;
  alert: AlertState;
}

interface CloseAlertAction {
  type: typeof CLOSE_ALERT;
}

export type AlertActions = OpenAlertAction | CloseAlertAction;

const alertState: AlertState = {
  title: "",
  content: "",
  onConfirm: null,
  open: false,
};

const alertReducer = (
  state: AlertState = alertState,
  action: AlertActions
): AlertState => {
  switch (action.type) {
    case OPEN_ALERT: {
      return {
        title: action.alert.title,
        content: action.alert.content,
        onConfirm: action.alert.onConfirm,
        open: true,
      };
    }
    case CLOSE_ALERT: {
      return { ...state, open: false };
    }
    default:
      return state;
  }
};

export default alertReducer;
