import { connect } from "react-redux";
const DisplayData = (props) => {
  const { testData } = props;

  return (
    <div>
      {testData.map((ele, index) => {
        return (
          <div key={index}>
            {ele.text}
            <br />
          </div>
        );
      })}
    </div>
  );
};

export default connect((state) => ({
  testData: state.testData,
}))(DisplayData);
