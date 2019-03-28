import { css } from 'styled-components';
import { math } from 'polished';

// Center horizontally and vertically
const centerHV = css`
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
`;

// Responsive square
const square = css`
	position: relative;
	border-radius: 50%;
	width: 100%;
	height: auto;
	padding-top: 100%;
`;

// Arrange children on a circle, last child in center
const onCircle = (
	count: number,
	circleSize: string,
	itemSize: string,
	angleShift = 0
) => css`
	position: relative;
	width: ${circleSize};
	height: ${circleSize};
	padding: 0;
	border-radius: 50%;
	list-style: none;

	> * {
		display: block;
		position: absolute;
		top: 50%;
		left: 50%;
		width: ${itemSize};
		height: ${itemSize};
		margin: ${math(`-${itemSize}/2`)};

		${() => {
			const angle = 360 / (count - 1);

			return [...Array(count - 1).keys()]
				.map(
					i => `
						&:nth-of-type(${i + 1}) {
							transform: rotate(${angle * i + angleShift}deg)
								translate(${math(`${circleSize} / 2`)}) rotate(-${angle * i + angleShift}deg);
						}
					`
				)
				.join();
		}}

		&:last-child {
			transform: unset;
		}
	}
`;

export { centerHV, square, onCircle };
