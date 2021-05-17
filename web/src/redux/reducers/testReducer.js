export default function testReducer(preState = [], action) {
  const { type, data } = action;
  switch (type) {
    case "getData":
      return [...data];
    default:
      return preState;
  }
}
