#ifndef TTLIBGO_UTIL_HPP
#define TTLIBGO_UTIL_HPP

#include <map>
#include <list>
#include <string>
using namespace std;

class maps {
public:
  map<string, string> strmap;
  map<string, uint32_t> uint32map;
  map<string, uint64_t> uint64map;
  map<string, list<string>> strlistmap;

  string getString(string key);
  uint32_t getUint32(string key);
  uint64_t getUint64(string key);
  list<string> getStringList(string key);
};

#endif