import React, { useEffect } from "react";
import Button from "@material-ui/core/Button";
import Switch from "@material-ui/core/Switch";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { Task } from "../Types";
import { useSelector, useDispatch } from "react-redux";
import { AppState } from "../reducer";
import { SingleTaskState, closeSingleTask } from "../reducer/task";
import { makeStyles } from "@material-ui/core/styles";
import MenuItem from "@material-ui/core/MenuItem";
import FormGroup from "@material-ui/core/FormGroup";
import FormControl from "@material-ui/core/FormControl";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Input from "@material-ui/core/Input";
import InputLabel from "@material-ui/core/InputLabel";
import Select from "@material-ui/core/Select";

const genHours = () => {
  let res: number[] = [];
  for (let i = 0; i < 24; i++) {
    res.push(i);
  }
  return res;
};

const genMinutes = () => {
  let res: number[] = [];
  for (let i = 0; i < 60; i++) {
    res.push(i);
  }
  return res;
};

const useStyles = makeStyles((theme) => ({
  root: {
    margin: theme.spacing(1),
    width: 300,
  },
  select: {
    margin: theme.spacing(1),
    width: 100,
  },
  group: {
    display: "block",
  },
}));

const TaskDialog: React.FC = () => {
  const task: SingleTaskState = useSelector<AppState, SingleTaskState>(
    (state) => state.singleTask
  );
  const [backend, setBackend] = React.useState("windows");
  const [address, setAddress] = React.useState("");
  const [user, setUser] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [enabled, setEnabled] = React.useState(true);
  const [sourcePath, setSourcePath] = React.useState("");
  const [destPath, setDestPath] = React.useState("");
  const [scheduleMinute, setScheduleMinute] = React.useState("0");
  const [scheduleHour, setScheduleHour] = React.useState("0");
  const dispatch = useDispatch();
  const classes = useStyles();
  const hours: number[] = genHours();
  const minutes: number[] = genMinutes();

  useEffect(() => {
    const curTask: Task | undefined = task.task;
    setBackend(curTask ? curTask.backend : "windows");
    setAddress(curTask ? curTask.address : "");
    setUser(curTask ? curTask.user : "");
    setPassword(curTask ? curTask.password : "");
    setEnabled(curTask ? curTask.enabled : true);
    setSourcePath(curTask ? curTask.sourcePath : "");
    setDestPath(curTask ? curTask.destPath : "");
    setScheduleMinute(curTask ? curTask.schedule.split(" ")[0] : "0");
    setScheduleHour(curTask ? curTask.schedule.split(" ")[1] : "0");
  }, [task]);

  const handleConfirm = () => {
    const newTask: Task = {
      id: task.task ? task.task.id : "",
      address: address,
      user: user,
      password: password,
      enabled: enabled,
      sourcePath: sourcePath,
      destPath: destPath,
      schedule: `${scheduleMinute} ${scheduleHour} * * *`,
      backend: backend,
    };
    task.onConfirm(newTask);
  };

  const handleClose = () => {
    dispatch(closeSingleTask());
  };

  const handleBackendChange = (event: any) => {
    setBackend(event.target.value);
  };
  const handleAddressChange = (event: any) => {
    setAddress(event.target.value);
  };
  const handleUserChange = (event: any) => {
    setUser(event.target.value);
  };
  const handlePasswordChange = (event: any) => {
    setPassword(event.target.value);
  };
  const handleEnabledChange = (event: any) => {
    setEnabled(event.target.checked);
  };
  const handleSourcePathChange = (event: any) => {
    setSourcePath(event.target.value);
  };
  const handleDestPathChange = (event: any) => {
    setDestPath(event.target.value);
  };
  const handleScheduleHourChange = (event: any) => {
    setScheduleHour(event.target.value);
  };
  const handleScheduleMinuteChange = (event: any) => {
    setScheduleMinute(event.target.value);
  };

  return (
    <div>
      <Dialog
        open={task.open}
        onClose={handleClose}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">
          {task.task ? "修改任务" : "新增任务"}
        </DialogTitle>
        <DialogContent>
          <form noValidate autoComplete="off">
            <FormControl className={classes.root}>
              <InputLabel htmlFor="task-address">地址</InputLabel>
              <Input
                id="task-address"
                value={address}
                onChange={handleAddressChange}
              />
            </FormControl>
            <FormControl className={classes.root}>
              <InputLabel htmlFor="task-user">用户名</InputLabel>
              <Input id="task-user" value={user} onChange={handleUserChange} />
            </FormControl>
            <FormControl className={classes.root}>
              <InputLabel htmlFor="task-password">密码</InputLabel>
              <Input
                id="task-password"
                value={password}
                onChange={handlePasswordChange}
                type="password"
              />
            </FormControl>
            <FormControl className={classes.root}>
              <InputLabel htmlFor="task-backend">操作系统</InputLabel>
              <Select
                id="task-backend"
                value={backend}
                onChange={handleBackendChange}
              >
                <MenuItem key="windows" value="windows">
                  Windows
                </MenuItem>
                <MenuItem key="linux" value="linux">
                  Linux
                </MenuItem>
              </Select>
            </FormControl>
            <FormGroup className={classes.group}>
              <FormControl className={classes.select}>
                <InputLabel htmlFor="task-schedule-hour">时</InputLabel>
                <Select
                  id="task-schedule-hour"
                  value={scheduleHour}
                  onChange={handleScheduleHourChange}
                >
                  {hours.map((v) => {
                    return (
                      <MenuItem key={v} value={v}>
                        {v}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
              <FormControl className={classes.select}>
                <InputLabel htmlFor="task-schedule-minute">分</InputLabel>
                <Select
                  id="task-schedule-minute"
                  value={scheduleMinute}
                  onChange={handleScheduleMinuteChange}
                >
                  {minutes.map((v) => {
                    return (
                      <MenuItem key={v} value={v}>
                        {v}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
              <FormControlLabel
                control={
                  <Switch
                    id="task-enabled"
                    checked={enabled}
                    onChange={handleEnabledChange}
                    color="primary"
                    inputProps={{ "aria-label": "primary checkbox" }}
                  />
                }
                className={classes.root}
                label={enabled ? "启用" : "禁用"}
              />
            </FormGroup>

            <FormControl className={classes.root}>
              <InputLabel htmlFor="task-source-path">源路径</InputLabel>
              <Input
                id="task-source-path"
                value={sourcePath}
                onChange={handleSourcePathChange}
              />
            </FormControl>
            <FormControl className={classes.root}>
              <InputLabel htmlFor="task-dest-path">目标路径</InputLabel>
              <Input
                id="task-dest-path"
                value={destPath}
                onChange={handleDestPathChange}
              />
            </FormControl>
          </form>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            取消
          </Button>
          <Button onClick={handleConfirm} color="primary">
            提交
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default TaskDialog;
