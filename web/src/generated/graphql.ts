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
  countriesAndTerritories?: Maybe<Scalars['String']>;
  cumulative_14d_per_10000: Scalars['String'];
};

export type Region = {
  __typename?: 'Region';
  dateRep: Scalars['String'];
  cases: Scalars['Int'];
  deaths: Scalars['Int'];
};

export type Query = {
  __typename?: 'Query';
  getRecords: Array<Record>;
  getRecordsByRegion: Array<Region>;
  getAllCountries: Array<Scalars['String']>;
};


export type QueryGetRecordsArgs = {
  countryName?: Maybe<Scalars['String']>;
};


export type QueryGetRecordsByRegionArgs = {
  region?: Maybe<Scalars['String']>;
};

