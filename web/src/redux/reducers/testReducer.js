export default function testReducer(
  preState = [
    {
      title: `Meeting room`,
      content: `Large high-level meeting room to improve work efficiency`,
      img: `https://spacestor.com/media/scaled_images/insights/amenity_rich/insights_spacestor_021219-072343_medium.jpg`,
    },
    {
      title: `Basketball court`,
      content: `Professional basketball court, instantly improve the level of competition`,
      img: `https://www.downtownmagazinenyc.com/wp-content/uploads/2018/03/SA13-Basketball_revG.jpg`,
    },
    {
      title: `Golf course`,
      content: `Upscale golf course, take you to experience high-end life`,
      img: `https://www.ciderridgegolf.com/images/slideshow/slide4.jpg`,
    },
  ],
  action
) {
  const { type, data } = action;
  switch (type) {
    case "getData":
      return [...data];
    default:
      return preState;
  }
}
