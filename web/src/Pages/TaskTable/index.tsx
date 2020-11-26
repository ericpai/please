import React, { useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import LinearProgress from "@material-ui/core/LinearProgress";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import InputAdornment from "@material-ui/core/InputAdornment";
import SearchIcon from "@material-ui/icons/Search";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import { useSnackbar } from "notistack";
import { Task, Response } from "../../Types";
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
  const [searchedValue, setSearchedValue] = useState<string>("");

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

  return (
    <React.Fragment>
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
          >
            添加
          </Button>
        </Grid>
      </Grid>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>地址</TableCell>
            <TableCell>系统</TableCell>
            <TableCell>启用状态</TableCell>
            <TableCell>上次执行结果</TableCell>
            <TableCell>上次执行时间</TableCell>
            <TableCell align="center">操作</TableCell>
            <TableCell align="center">详细信息</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {visibleTasks &&
            visibleTasks.map((task: Task) => (
              <Row key={task.id} task={task} onNotify={onNotify}></Row>
            ))}
        </TableBody>
      </Table>
    </React.Fragment>
  );
};

export default TaskTable;