import React, { useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import LinearProgress from "@material-ui/core/LinearProgress";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TextField from "@material-ui/core/TextField";
import TablePagination from "@material-ui/core/TablePagination";
import TableContainer from "@material-ui/core/TableContainer";
import Button from "@material-ui/core/Button";
import InputAdornment from "@material-ui/core/InputAdornment";
import SearchIcon from "@material-ui/icons/Search";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import { useSnackbar } from "notistack";
import { Task, Request, Response } from "../../Types";
import TaskDialog from "../../Components/TaskDialog";
import { useDispatch } from "react-redux";
import {
  SingleTaskState,
  openSingleTask,
  closeSingleTask,
} from "../../reducer/task";
import Row from "./Row";

const useStyles = makeStyles((theme) => ({
  root: {
    marginBottom: theme.spacing(2),
    marginTop: theme.spacing(2),
  },
  newBtn: {
    float: "right",
  },
}));

const TaskTable: React.FC = () => {
  const { enqueueSnackbar } = useSnackbar();
  const classes = useStyles();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [visibleTasks, setVisibleTasks] = useState<Task[]>([]);
  const [page, setPage] = React.useState(0);
  const dispatch = useDispatch();
  const [rowsPerPage, setRowsPerPage] = React.useState(10);

  useEffect(() => {
    if (!loading) {
      return;
    }
    fetch(`/api/tasks`)
      .then((response) => {
        if (!response.ok) {
          throw new Error("请求失败");
        }
        return response.json();
      })
      .then((body: Response) => {
        const tasks = body.data.tasks;
        setTasks(tasks);
        setVisibleTasks(tasks);
      })
      .catch((reason) => {
        enqueueSnackbar(`请求失败：${reason.toString()}`, {
          variant: "error",
        });
      })
      .finally(() => {
        setLoading(false);
      });
  }, [loading]);

  const onNotify = () => {
    setLoading(true);
  };

  const handleChangePage = (event: any, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: any) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const handleSearchValueChange = (event: any) => {
    let newTasks: Task[] = [];
    for (let i = 0; i < tasks.length; i++) {
      if (tasks[i].address.includes(event.target.value)) {
        newTasks.push(tasks[i]);
      }
    }
    setVisibleTasks(newTasks);
    setPage(0);
  };

  const onCreate = () => {
    const newState: SingleTaskState = {
      task: undefined,
      open: true,
      onConfirm: (t: Task) => {
        handleCreate(t);
      },
    };
    dispatch(openSingleTask(newState));
  };

  const handleCreate = (task: Task) => {
    const req: Request = {
      task: task,
    };
    fetch(`/api/tasks`, {
      method: "POST",
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
          enqueueSnackbar(`添加成功`, {
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
      <TaskDialog />
      <Typography component="h2" variant="h6" color="primary" gutterBottom>
        任务列表
      </Typography>
      {loading ? <LinearProgress /> : <Divider></Divider>}
      <Grid container className={classes.root}>
        <Grid item xs>
          <TextField
            id="outlined-search"
            label="搜索 IP"
            type="search"
            variant="outlined"
            size="small"
            fullWidth
            onChange={handleSearchValueChange}
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <SearchIcon />
                </InputAdornment>
              ),
            }}
          />
        </Grid>
        <Grid item xs={6}></Grid>
        <Grid item xs>
          <Button
            variant="contained"
            color="primary"
            className={classes.newBtn}
            onClick={onCreate}
          >
            添加
          </Button>
        </Grid>
      </Grid>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>地址</TableCell>
              <TableCell>系统</TableCell>
              <TableCell>启用状态</TableCell>
              <TableCell>执行计划</TableCell>
              <TableCell>上次执行结果</TableCell>
              <TableCell>上次执行时间</TableCell>
              <TableCell align="center">操作</TableCell>
              <TableCell align="center">详细信息</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {visibleTasks &&
              visibleTasks
                .slice(rowsPerPage * page, rowsPerPage * (page + 1))
                .map((task: Task) => (
                  <Row key={task.id} task={task} onNotify={onNotify}></Row>
                ))}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10, 20, 50]}
        colSpan={12}
        count={visibleTasks ? visibleTasks.length : 0}
        rowsPerPage={rowsPerPage}
        page={page}
        component="div"
        onChangePage={handleChangePage}
        onChangeRowsPerPage={handleChangeRowsPerPage}
      />
    </React.Fragment>
  );
};

export default TaskTable;
