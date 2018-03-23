import React from 'react';
import { mount, shallow } from 'enzyme';
import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
Enzyme.configure({ adapter: new Adapter() });

import Home from './Home';

describe('App', () => {
  const mockFn = jest.fn()


  it('App Mounts without Crashing', () => {
    const Container = mount(<Home/>)
  })
})