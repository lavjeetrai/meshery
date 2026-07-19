import { describe, expect, it, vi } from 'vitest';
import { generatePerformanceProfile } from './performance-helpers';

vi.mock('./helper', () => ({
  generateTestName: vi.fn((name, serviceMesh) => `mocked-name-${name}-${serviceMesh}`),
}));

describe('generatePerformanceProfile', () => {
  it('generates a performance profile object correctly with an id', () => {
    const inputData = {
      id: 'profile-123',
      name: 'test-profile',
      loadGenerator: 'fortio',
      additional_options: { flag: 'value' },
      endpoint: 'http://localhost',
      serviceMesh: 'istio',
      concurrentRequest: 10,
      qps: 100,
      duration: '30s',
      requestHeaders: 'header1: value1',
      requestCookies: 'cookie1=value1',
      requestBody: '{"key": "value"}',
      contentType: 'application/json',
      caCertificate: { file: 'cert-content', name: 'my-cert' },
    };

    const result = generatePerformanceProfile(inputData);

    expect(result).toEqual({
      id: 'profile-123',
      name: 'mocked-name-test-profile-istio',
      loadGenerators: ['fortio'],
      endpoints: ['http://localhost'],
      serviceMesh: 'istio',
      concurrentRequest: 10,
      qps: 100,
      duration: '30s',
      requestHeaders: 'header1: value1',
      requestBody: '{"key": "value"}',
      requestCookies: 'cookie1=value1',
      contentType: 'application/json',
      metadata: {
        additional_options: [{ flag: 'value' }],
        ca_certificate: {
          file: 'cert-content',
          name: 'my-cert',
        },
      },
    });
  });

  it('generates a performance profile object correctly without an id', () => {
    const inputData = {
      name: 'test-profile',
      loadGenerator: 'fortio',
      additional_options: { flag: 'value' },
      endpoint: 'http://localhost',
      serviceMesh: 'istio',
      concurrentRequest: 10,
      qps: 100,
      duration: '30s',
      requestHeaders: 'header1: value1',
      requestCookies: 'cookie1=value1',
      requestBody: '{"key": "value"}',
      contentType: 'application/json',
      caCertificate: { file: 'cert-content', name: 'my-cert' },
    };

    const result = generatePerformanceProfile(inputData);

    expect(result).toEqual({
      name: 'mocked-name-test-profile-istio',
      loadGenerators: ['fortio'],
      endpoints: ['http://localhost'],
      serviceMesh: 'istio',
      concurrentRequest: 10,
      qps: 100,
      duration: '30s',
      requestHeaders: 'header1: value1',
      requestBody: '{"key": "value"}',
      requestCookies: 'cookie1=value1',
      contentType: 'application/json',
      metadata: {
        additional_options: [{ flag: 'value' }],
        ca_certificate: {
          file: 'cert-content',
          name: 'my-cert',
        },
      },
    });
  });
});
