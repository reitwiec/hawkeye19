import React, { Component } from 'react';
import styled from 'styled-components';

import { Button, TextField } from '../../components';

type Props = {
	className: string;
};

type State = {
	formData: {
		[name: string]: string;
		region: string;
		level: string;
		question: string;
		answer: string;
		addinfo: string;
		addedBy: string;
	};
};

class QuestionPage extends Component<Props, State> {
	state = {
		formData: {
			region: '',
			level: '',
			question: '',
			answer: '',
			addinfo: '',
			addedBy: ''
		}
	};

	onChange = (name, value) => {
		this.setState({
			formData: { ...this.state.formData, [name]: value }
		});
	};

	submitQuestion = () => {
		const formData = {
			...this.state.formData,
			region: parseInt(this.state.formData.region),
			level: parseInt(this.state.formData.level)
		};
		// console.log(formData);
		fetch('/api/addQuestion', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(formData)
		})
			.then(res => res.json())
			.then(json => console.log(json));
	};

	render() {
		return (
			<div className={this.props.className}>
				<h2>Questions</h2>
				<div className="question-form">
					<div className="region-level">
						<TextField
							name="region"
							placeholder="Region"
							onChange={this.onChange}
						/>
						<TextField
							name="level"
							placeholder="Level"
							onChange={this.onChange}
						/>
					</div>
					<TextField
						name="question"
						placeholder="Question"
						onChange={this.onChange}
					/>
					<TextField
						name="answer"
						placeholder="Answer"
						onChange={this.onChange}
					/>
					<TextField
						name="addinfo"
						placeholder="Add Info"
						onChange={this.onChange}
					/>
					<TextField
						name="addedBy"
						placeholder="Added By"
						onChange={this.onChange}
					/>
					<Button onClick={this.submitQuestion}>Add</Button>
				</div>
				<div className="questions-container">{}</div>
			</div>
		);
	}
}

export default styled(QuestionPage)`
	.question-form {
		display: flex;
		flex-flow: column nowrap;
		width: min-content;

		.region-level {
			display: flex;

			${TextField} {
				margin-right: 0.5em;
				input {
					width: 9ch;
				}
			}
		}
	}
`;
