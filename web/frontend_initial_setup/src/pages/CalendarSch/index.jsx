import { connect } from "react-redux";
import { getData } from "../../redux/actions/testAction";
import DisplayData from "./DisplayData";
const CalendarSch = (props) => {
  const { getData } = props;
  const clickHandler = async () => {
    try {
      const response = await fetch("https://cat-fact.herokuapp.com/facts");
      if (response.ok) {
        const data = await response.json();
        getData(data);
      }
    } catch (e) {
      console.log(e);
    }
  };
  return (
    <div>
      this is Calendar:
      <br />
      <button onClick={clickHandler}>click me</button>
      <br />
      <DisplayData />
    </div>
  );
};

export default connect(
  (state) => ({
    testData: state.testData,
  }),
  { getData }
)(CalendarSch);
