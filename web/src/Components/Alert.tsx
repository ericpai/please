import React from "react";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogTitle from "@material-ui/core/DialogTitle";
import { useSelector, useDispatch } from "react-redux";
import { AppState } from "../reducer";
import { AlertState, closeAlert } from "../reducer/alert";

const AlertDialog: React.FC = () => {
  const alert: AlertState = useSelector<AppState, AlertState>(
    (state) => state.alert
  );
  const dispatch = useDispatch();

  const handleConfirm = () => {
    alert.onConfirm();
    dispatch(closeAlert());
  };

  const handleClose = () => {
    dispatch(closeAlert());
  };

  return (
    <div>
      <Dialog
        open={alert.open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{alert.title}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            {alert.content}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary" autoFocus>
            取消
          </Button>
          <Button onClick={handleConfirm} color="primary" autoFocus>
            确定
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default AlertDialog;
