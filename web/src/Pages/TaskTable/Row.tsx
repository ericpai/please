import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { red, green } from "@material-ui/core/colors";
import CheckIcon from "@material-ui/icons/Check";
import ClearIcon from "@material-ui/icons/Clear";
import PlayArrow from "@material-ui/icons/PlayArrow";
import KeyboardArrowDownIcon from "@material-ui/icons/KeyboardArrowDown";
import KeyboardArrowUpIcon from "@material-ui/icons/KeyboardArrowUp";
import DeleteIcon from "@material-ui/icons/Delete";
import EditIcon from "@material-ui/icons/Edit";
import { Icon } from "@iconify/react";
import linuxIcon from "@iconify-icons/mdi/linux";
import microsoftWindows from "@iconify-icons/mdi/microsoft-windows";
import Stop from "@material-ui/icons/Stop";
import Paper from "@material-ui/core/Paper";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import Collapse from "@material-ui/core/Collapse";
import Typography from "@material-ui/core/Typography";
import Box from "@material-ui/core/Box";
import IconButton from "@material-ui/core/IconButton";
import { useDispatch } from "react-redux";
import { AlertState, openAlert } from "../../reducer/alert";
import {
  SingleTaskState,
  openSingleTask,
  closeSingleTask,
} from "../../reducer/task";
import { Task, Request, Response } from "../../Types";
import { humanizedDatetime } from "../../Utils";
import { useSnackbar } from "notistack";

const useStyles = makeStyles({
  root: {
    "& > *": {
      borderBottom: "unset",
    },
  },
  success: {
    color: green[700],
  },
  error: {
    color: red[700],
  },
});

interface Props {
  task: Task;
  onNotify: any;
}

const Row: React.FC<Props> = (props: Props) => {
  const classes = useStyles();
  const task: Task = props.task;
  const onNotify: any = props.onNotify;
  const [open, setOpen] = React.useState(false);
  const dispatch = useDispatch();
  const { enqueueSnackbar } = useSnackbar();

  const formatSchedule = (sche: string) => {
    const parts = sche.split(" ");
    return `${parts[1]} 时 ${parts[0]} 分`;
  };

  const onDelete = (taskId: string) => {
    const newState: AlertState = {
      title: "删除任务",
      content: "确定要删除吗？(被删除的任务无法恢复)",
      onConfirm: () => {
        handleDelete(taskId);
      },
      open: true,
    };
    dispatch(openAlert(newState));
  };
  const handleDelete = (taskId: string) => {
    fetch(`/api/tasks/${taskId}`, {
      method: "DELETE",
    })
      .then((response) => {
        if (response.status !== 204) {
          throw new Error("请求失败");
        }
        enqueueSnackbar(`删除成功`, {
          variant: "success",
        });
      })
      .catch((reason) => {
        enqueueSnackbar(`请求失败：${reason.toString()}`, {
          variant: "error",
        });
      })
      .finally(() => {
        onNotify();
      });
  };

  const onUpdate = (task: Task) => {
    const newState: SingleTaskState = {
      task: task,
      open: true,
      onConfirm: (t: Task) => {
        handleUpdate(t);
      },
    };
    dispatch(openSingleTask(newState));
  };

  const handleUpdate = (task: Task) => {
    const req: Request = {
      task: task,
    };
    fetch(`/api/tasks/${task.id}`, {
      method: "PATCH",
      body: JSON.stringify(req),
    })
      .then((response) => {
        return response.json();
      })
      .then((resp: Response) => {
        if (resp.meta.code !== 201) {
          enqueueSnackbar(resp.meta.message, {
            variant: "error",
          });
        } else {
          enqueueSnackbar(`修改成功`, {
            variant: "success",
          });
          dispatch(closeSingleTask());
          onNotify();
        }
      })
      .catch((reason) => {
        enqueueSnackbar(`请求失败：${reason.toString()}`, {
          variant: "error",
        });
      });
  };
  return (
    <React.Fragment>
      <TableRow key={task.id} className={classes.root}>
        <TableCell>{task.id}</TableCell>
        <TableCell>{task.address}</TableCell>
        <TableCell>
          <Icon
            icon={task.backend === "windows" ? microsoftWindows : linuxIcon}
            width={24}
            height={24}
          />
        </TableCell>
        <TableCell>
          {task.enabled ? (
            <Stop className={classes.error}></Stop>
          ) : (
            <PlayArrow className={classes.success}></PlayArrow>
          )}
        </TableCell>
        <TableCell>{formatSchedule(task.schedule)}</TableCell>
        <TableCell>
          {task.succeed ? (
            <ClearIcon className={classes.error}></ClearIcon>
          ) : (
            <CheckIcon className={classes.success}></CheckIcon>
          )}
        </TableCell>
        <TableCell>
          {humanizedDatetime(task.updatedTime ? task.updatedTime : 0)}
        </TableCell>
        <TableCell align="center">
          <IconButton
            aria-label="edit"
            className={classes.success}
            onClick={() => onUpdate(task)}
          >
            <EditIcon />
          </IconButton>
          <IconButton
            aria-label="delete"
            className={classes.error}
            onClick={() => onDelete(task.id)}
          >
            <DeleteIcon />
          </IconButton>
        </TableCell>
        <TableCell align="center">
          <IconButton
            aria-label="expand row"
            size="small"
            onClick={() => setOpen(!open)}
          >
            {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
          </IconButton>
        </TableCell>
      </TableRow>
      <TableRow>
        <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={9}>
          <Collapse in={open} timeout="auto" unmountOnExit>
            <Box margin={1}>
              <Paper elevation={0}>
                <Typography variant="h6" gutterBottom component="div">
                  详细信息
                </Typography>
                <Typography variant="body1" component="p">
                  源路径：{task.sourcePath}
                </Typography>
                <Typography variant="body1" component="p">
                  目标路径：{task.destPath}
                </Typography>
                <Typography variant="body1" component="p">
                  创建时间：
                  {humanizedDatetime(task.createdTime ? task.createdTime : 0)}
                </Typography>
              </Paper>
            </Box>
          </Collapse>
        </TableCell>
      </TableRow>
    </React.Fragment>
  );
};

export default Row;
