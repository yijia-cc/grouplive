import { Route, Switch, Redirect } from "react-router-dom";
import { connect } from "react-redux";
import { getData } from "../../redux/actions/testAction";
import SubNavBar from "./SubNavBar";
import AmenityList from "./AmenityList";
import History from "./History";
import CalendarSchedule from "./CalendarSchedule";
import "./index.css";

const Calendar = (props) => {
  const { testData } = props;
  const { location, history } = props;
  const amenityList = () => (
    <AmenityList amenityList={testData} history={history} />
  );
  return (
    <div>
      <SubNavBar location={location} />
      <div className="calendar-main">
        <Switch>
          <Route path="/calendar/amenitylist" render={amenityList} />
          <Route path="/calendar/history" component={History} />
          <Route
            path="/calendar/calendarScheduler"
            component={CalendarSchedule}
          />
          <Redirect to="/calendar/amenitylist" />
        </Switch>
      </div>
    </div>
  );
};

export default connect(
  (state) => ({
    testData: state.testData,
  }),
  { getData }
)(Calendar);
