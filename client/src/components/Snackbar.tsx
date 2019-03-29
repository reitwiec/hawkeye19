import React, { Component } from 'react';
import styled, { css } from 'styled-components';

interface ISnackbarProps {
	message: JSX.Element | string;
	action: JSX.Element | null;
	open: boolean;
	autoHideDelay: number;
}

class Snackbar extends Component<ISnackbarProps> {
	public static defaultProps = {
		action: null,
		autoHideDelay: 0,
		open: false,
	};

	public render() {
		return (
			<SnackbarWrapper {...this.props}>{this.props.message}</SnackbarWrapper>
		);
	}
}

const SnackbarWrapper = styled.div`
	display: block;
	position: fixed;
	bottom: 0;
	width: 100vw;
	background-color: #555;
	padding: 0.5em;
	color: #bbbbee;
	font-family: sans-serif;
	text-transform: uppercase;
	opacity: 0;
	transition: all 0.25s ease;
	margin-bottom: -2em;
	z-index: 999;

	${(props: ISnackbarProps) =>
		props.open
			? css`
					opacity: 1;
					margin-bottom: 0;
			  `
			: ``}
`;

export default Snackbar;
