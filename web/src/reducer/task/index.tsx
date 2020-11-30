import { Task } from "../../Types";

export interface SingleTaskState {
  task?: Task;
  onConfirm: any;
  open: boolean;
}

export const OPEN_SINGLE_TASK = "OPEN_SINGLE_TASK";
export const CLOSE_SINGLE_TASK = "CLOSE_SINGLE_TASK";

export function openSingleTask(task: SingleTaskState): SingleTaskActions {
  return {
    type: OPEN_SINGLE_TASK,
    task: task,
  };
}

export function closeSingleTask(): SingleTaskActions {
  return {
    type: CLOSE_SINGLE_TASK,
  };
}

interface OpenSingleTaskAction {
  type: typeof OPEN_SINGLE_TASK;
  task: SingleTaskState;
}

interface CloseSingleTaskAction {
  type: typeof CLOSE_SINGLE_TASK;
}

export type SingleTaskActions = OpenSingleTaskAction | CloseSingleTaskAction;

const singleTaskState: SingleTaskState = {
  task: undefined,
  onConfirm: null,
  open: false,
};

const singleTaskReducer = (
  state: SingleTaskState = singleTaskState,
  action: SingleTaskActions
): SingleTaskState => {
  switch (action.type) {
    case OPEN_SINGLE_TASK: {
      return {
        task: action.task.task,
        onConfirm: action.task.onConfirm,
        open: true,
      };
    }
    case CLOSE_SINGLE_TASK: {
      return { ...state, open: false };
    }
    default:
      return state;
  }
};

export default singleTaskReducer;
