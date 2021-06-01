import React from "react";
import Paper from "@material-ui/core/Paper";
import { ViewState, EditingState } from "@devexpress/dx-react-scheduler";
import {
  Scheduler,
  Appointments,
  AppointmentForm,
  AppointmentTooltip,
  WeekView,
  ConfirmationDialog,
  CurrentTimeIndicator,
} from "@devexpress/dx-react-scheduler-material-ui";
import { appointments } from "./demo-data/appointments";
import CalendarHeader from "./CalendarHeader";
import "./index.css";

export default class CalendarSchedule extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      data: appointments,
      currentDate: Date.now(),

      addedAppointment: {},
      appointmentChanges: {},
      editingAppointment: undefined,
    };

    this.commitChanges = this.commitChanges.bind(this);
    this.changeAddedAppointment = this.changeAddedAppointment.bind(this);
    this.changeAppointmentChanges = this.changeAppointmentChanges.bind(this);
    this.changeEditingAppointment = this.changeEditingAppointment.bind(this);
    this.currentDateChange = (currentDate) => {
      this.setState({ currentDate });
    };
  }

  changeAddedAppointment(addedAppointment) {
    this.setState({ addedAppointment });
  }

  changeAppointmentChanges(appointmentChanges) {
    this.setState({ appointmentChanges });
  }

  changeEditingAppointment(editingAppointment) {
    this.setState({ editingAppointment });
  }

  commitChanges({ added, changed, deleted }) {
    this.setState((state) => {
      let { data } = state;
      if (added) {
        const startingAddedId =
          data.length > 0 ? data[data.length - 1].id + 1 : 0;
        data = [...data, { id: startingAddedId, ...added }];
      }
      if (changed) {
        data = data.map((appointment) =>
          changed[appointment.id]
            ? { ...appointment, ...changed[appointment.id] }
            : appointment
        );
      }
      if (deleted !== undefined) {
        data = data.filter((appointment) => appointment.id !== deleted);
      }
      return { data };
    });
  }

  timeComparatoCurrentTime = (time1) =>
    time1 < this.state.currentDate ? true : false;
  timeTableCellComponent = ({ ...restProps }) => {
    const time1 = restProps.startDate.getTime();
    const checkIfTimePast = this.timeComparatoCurrentTime(time1);
    if (checkIfTimePast) {
      restProps.onDoubleClick = () => {
        console.log("time has Past");
      };
    }
    return <WeekView.TimeTableCell {...restProps} />;
  };

  headerComponent = ({ showOpenButton, showDeleteButton, ...restProps }) => {
    const { startDate } = restProps.appointmentData;
    const checkIfTimePast = this.timeComparatoCurrentTime(startDate.getTime());
    return (
      <AppointmentTooltip.Header
        showOpenButton={!checkIfTimePast}
        showDeleteButton={!checkIfTimePast}
        {...restProps}
      />
    );
  };

  render() {
    const { state } = this.props.history.location;
    const {
      currentDate,
      data,
      addedAppointment,
      appointmentChanges,
      editingAppointment,
    } = this.state;

    return (
      <>
        <CalendarHeader state={state} />
        <Paper className="calendar-body" state={state}>
          <Scheduler data={data} height={760}>
            <ViewState currentDate={currentDate} />
            <EditingState
              onCommitChanges={this.commitChanges}
              addedAppointment={addedAppointment}
              onAddedAppointmentChange={this.changeAddedAppointment}
              appointmentChanges={appointmentChanges}
              onAppointmentChangesChange={this.changeAppointmentChanges}
              editingAppointment={editingAppointment}
              onEditingAppointmentChange={this.changeEditingAppointment}
            />
            <WeekView
              startDayHour={9}
              endDayHour={22}
              timeTableCellComponent={this.timeTableCellComponent}
            />
            <ConfirmationDialog />
            <Appointments />
            <AppointmentTooltip
              headerComponent={this.headerComponent}
              showOpenButton
              showDeleteButton
            />
            <AppointmentForm />
            <CurrentTimeIndicator
              shadePreviousCells={true}
              shadePreviousAppointments={true}
              updateInterval={10000}
            />
          </Scheduler>
        </Paper>
      </>
    );
  }
}
