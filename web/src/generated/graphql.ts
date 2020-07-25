export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: any }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Record = {
  __typename?: 'Record';
  id: Scalars['ID'];
  dateRep: Scalars['String'];
  cases: Scalars['Int'];
  deaths: Scalars['Int'];
  countriesAndTerritories: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  recordsByCountryName: Array<Record>;
  getAllCountries: Array<Scalars['String']>;
};


export type QueryRecordsByCountryNameArgs = {
  countryName: Scalars['String'];
};

