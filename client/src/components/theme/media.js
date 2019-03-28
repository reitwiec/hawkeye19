import { css } from 'styled-components';
const sizes = {
	phone: 420,
	tablet: 768,
	desktop: 992,
	giant: 1170
};

const query = size => (...args) => {
	return (...args) => {
		return css`
			@media (max-width: ${size}px) {
				${css(...args)}
			}
		`;
	};
};

const desktop = query(sizes.desktop);
const giant = query(sizes.giant);

function phone(...args) {
	return css`
		@media (max-width: ${sizes.phone}px) {
			${css(...args)}
		}
	`;
}

function tablet(...args) {
	return css`
		@media (max-width: ${sizes.tablet}px) {
			${css(...args)}
		}
	`;
}

const media = {
	phone,
	tablet,
	desktop,
	giant,
	query
};

export default media;
