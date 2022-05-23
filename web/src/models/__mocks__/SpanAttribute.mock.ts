import faker from '@faker-js/faker';
import {IMockFactory} from '../../types/Common.types';
import {TSpanFlatAttribute} from '../../types/Span.types';
import {TSpanAttribute} from '../../types/SpanAttribute.types';
import SpanAttribute from '../SpanAttribute.model';

const SpanAttributeMock: IMockFactory<TSpanAttribute, TSpanFlatAttribute> = () => ({
  raw(data = {}) {
    return {
      value: faker.random.word(),
      key: faker.random.word(),
      ...data,
    };
  },
  model(data = {}) {
    return SpanAttribute(this.raw(data));
  },
});

export default SpanAttributeMock();
