import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";
import media from "../components/theme/media";

const QuestionPage = ({ className, ...props }) => {
  const check = () => {
    console.log("hello");
  };
  const { region } = props.location.state || { region: "Installation 09" };
  return (
    <div className={className}>
      <p>{region}</p>
      <p>Level:24</p>
      <div id="question">
        Lorem ipsum, dolor sit amet consectetur adipisicing elit. Deserunt
        debitis est, vel accusamus aliquam eos pariatur laborum ratione ullam
        quae ipsam ad ipsa, necessitatibus placeat assumenda veritatis maxime
        excepturi repellat?
      </div>
      <textarea
        className="form__input"
        id="answer"
        type="text"
        placeholder="Write the answer here"
        aria-invalid="false"
      />

      <div>
        <button onClick={check} id="submit">
          Submit
        </button>
      </div>

      <div id="hints">
        <p>Hints:</p>
        Muh mein le le
      </div>
    </div>
  );
};

QuestionPage.propTypes = {
  className: PropTypes.string
};

export default styled(QuestionPage)`
  ${media.phone`
	color:red;
	#question{
		background:blue;
		text-align:center;
		height:300px;
		margin: 10px;
		padding:20px;
		font-size:1.2em;
	}
	#answer{
		background:blue;
		text-align:center;
		height:auto;
		width:335px;
		margin: 10px;
		padding:20px;
		font-size:1.2em;
		color:red;
		border:none;
	}
	#hints{
		background:blue;
		text-align:center;
		height:150px;
		margin: 10px;
		padding:5px;
		font-size:1.2em;
		color:red;}
	#submit{
		 background:red;
		 margin-left:50%;
		 margin-right:50%;
	}
`}
`;
