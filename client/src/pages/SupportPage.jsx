import React, { Component } from "react";
import styled from "styled-components";

class SupportPage extends Component {
  render() {
    return (
      <div className={this.props.className}>
        <div class="form__group">
          <h1>Send us your queries!</h1>
          <textarea
            rows="7"
            class="form__input"
            id="textarea-disabled"
            type="text"
            placeholder="Write to us here"
            aria-invalid="false"
          />
        </div>
        <button>Submit</button>
      </div>
    );
  }
}

export default styled(SupportPage)`
  textarea {
    resize: none;
    width: 600px;
    font-size: 1.2em;
  }
`;
